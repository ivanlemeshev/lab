package gateway

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"

	auth "lem/go/auth/internal/auth"
)

// AuthenticationMiddleware intercepts HTTP requests
func AuthenticationMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Gateway: Received request for %s", r.URL.Path)

		// 1. Skip auth for the public login endpoint
		if r.URL.Path == "/v1/login" {
			h.ServeHTTP(w, r)
			return
		}

		// 2. Find the token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// 3. Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			log.Printf("Token validation error: %v", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// 4. Token is valid!
		// We add the user's identity as *new* HTTP headers.
		// The gateway (in main.go) will be configured to turn
		// these into gRPC metadata.
		r.Header.Set("X-User-Id", claims.UserID)
		r.Header.Set("X-User-Role", claims.Role)

		// We've authenticated. Pass the request to the gRPC-Gateway mux.
		h.ServeHTTP(w, r)
	})
}

// CustomMetadataAnnotator is the function that the gRPC-Gateway will use
// to map our HTTP headers (like X-User-Id) to gRPC metadata.
func CustomMetadataAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}
	if userID := req.Header.Get("X-User-Id"); userID != "" {
		md.Set("x-user-id", userID)
	}
	if userRole := req.Header.Get("X-User-Role"); userRole != "" {
		md.Set("x-user-role", userRole)
	}
	return md
}

// NewGatewayMux creates and configures the gRPC-Gateway multiplexer
func NewGatewayMux() *runtime.ServeMux {
	// This is the key! We tell the gateway to use our
	// CustomMetadataAnnotator to map headers to gRPC metadata.
	return runtime.NewServeMux(
		runtime.WithMetadata(CustomMetadataAnnotator),
	)
}
