package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	users "lem/go/auth/internal/users"
)

var jwtSecret = []byte("my-super-secret-key")

// Claims defines the structure for JWT claims.
type Claims struct {
	jwt.RegisteredClaims
	UserID string     `json:"user_id"`
	Role   users.Role `json:"role"`
}

// GenerateToken creates a JWT token for a given user ID and role.
func GenerateToken(userID string, role users.Role) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 1)), // 1 hour expiry
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}

// ValidateToken verifies the JWT token and returns the claims if valid.
func ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
