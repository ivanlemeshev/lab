package services

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "lem/go/auth/proto"
)

// OrdersServer implements the OrdersService
type OrdersServer struct {
	pb.UnimplementedOrdersServiceServer
}

// GetOrder is a protected endpoint
func (s *OrdersServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	// Read identity from gRPC metadata
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

	// **AUTHORIZATION LOGIC**
	// Only admins can access orders.
	if userRole[0] != "admin" {
		return nil, status.Error(codes.PermissionDenied, "only admins can access orders")
	}

	// Dummy response
	return &pb.OrderResponse{
		OrderId: req.OrderId,
		Item:    "Dummy Item",
		UserId:  "u-user-123", // The user who placed the order
	}, nil
}
