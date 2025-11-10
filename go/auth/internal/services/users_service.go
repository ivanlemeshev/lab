package services

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	repos "lem/go/auth/internal/repos"
	pb "lem/go/auth/proto"
)

// UsersServer implements the UsersService.
type UsersServer struct {
	pb.UnimplementedUsersServiceServer
	users repos.UserRepository
}

// NewUsersServer creates a new UsersServer.
func NewUsersServer(usersRepo repos.UserRepository) *UsersServer {
	return &UsersServer{
		users: usersRepo,
	}
}

// GetUser retrieves user information.
func (s *UsersServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
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

	if userRole[0] != "admin" && userID[0] != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "you can only access your own user data")
	}

	user, err := s.users.FindByUsername(req.UserId)
	if errors.Is(err, repos.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.UserResponse{
		UserId: user.ID,
		Name:   user.Username,
		Role:   user.Role.String(),
	}, nil
}
