package services

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	users "lem/go/auth/internal/users"
	pb "lem/go/auth/proto"
)

// OrdersServer implements the OrdersService
type OrdersServer struct {
	pb.UnimplementedOrdersServiceServer
}

// NewOrdersServer creates a new OrdersServer.
func NewOrdersServer() *OrdersServer {
	return &OrdersServer{}
}

// GetOrder retrieves order information.
func (s *OrdersServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	userID := md.Get("x-user-id")
	userRole := md.Get("x-user-role")

	if len(userID) == 0 || len(userRole) == 0 {
		return nil, status.Error(codes.Unauthenticated, "user identity not in metadata")
	}

	log.Printf("GetOrder called by UserID: %s, Role: %s", userID[0], userRole[0])

	if userRole[0] != users.RoleAdmin.String() {
		return nil, status.Error(codes.PermissionDenied, "only admins can access orders")
	}

	return &pb.OrderResponse{
		OrderId: req.OrderId,
		Item:    "Dummy Item",
		UserId:  "u-user-123",
	}, nil
}
