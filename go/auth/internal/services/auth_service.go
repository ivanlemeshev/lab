package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth "lem/go/auth/internal/auth"
	pb "lem/go/auth/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var role string
	var userID string

	if req.Username == "admin" && req.Password == "admin" {
		userID = "u-admin-001"
		role = "admin"
	} else if req.Username == "user" && req.Password == "user" {
		userID = "u-user-123"
		role = "user"
	} else {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	token, err := auth.GenerateToken(userID, role)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate token")
	}

	return &pb.LoginResponse{Token: token}, nil
}
