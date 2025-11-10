package services

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth "lem/go/auth/internal/auth"
	repos "lem/go/auth/internal/repos"
	pb "lem/go/auth/proto"
)

// AuthServer implements the AuthService.
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	users repos.UserRepository
}

// NewAuthServer creates a new AuthServer.
func NewAuthServer(usersRepo repos.UserRepository) *AuthServer {
	return &AuthServer{
		users: usersRepo,
	}
}

// Login authenticates a user and returns a JWT token.
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.users.FindByUsername(req.Username)
	if errors.Is(err, repos.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user.Password != req.Password {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate token")
	}

	return &pb.LoginResponse{Token: token}, nil
}
