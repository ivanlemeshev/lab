package services

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "lem/go/auth/proto"
)

// UsersServer implements the UsersService
type UsersServer struct {
	pb.UnimplementedUsersServiceServer
}

// GetUser is a protected endpoint
func (s *UsersServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// **THIS IS THE KEY PART**
	// We read the identity from the gRPC metadata passed by the gateway
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	userID := md.Get("x-user-id")
	userRole := md.Get("x-user-role")

	if len(userID) == 0 || len(userRole) == 0 {
		return nil, status.Error(codes.Unauthenticated, "user identity not in metadata")
	}

	log.Printf("GetUser called by UserID: %s, Role: %s", userID[0], userRole[0])

	// **AUTHORIZATION LOGIC**
	// A user can only get their own info, but an admin can get anyone's info.
	if userRole[0] != "admin" && userID[0] != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "you can only access your own user data")
	}

	// Dummy response (in a real app, query a database)
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   "Dummy User Name",
		Role:   "user", // This would be from the DB
	}, nil
}
