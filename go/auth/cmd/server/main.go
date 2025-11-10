package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gateway "lem/go/auth/internal/gateway"
	repos "lem/go/auth/internal/repos"
	services "lem/go/auth/internal/services"
	pb "lem/go/auth/proto"
)

const (
	grpcServerAddr = "localhost:50051"
	httpServerAddr = "localhost:8080"
)

func startGRPCServer() {
	lis, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	usersRepo := repos.NewFakeUsersRepository()

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, services.NewAuthServer(usersRepo))
	pb.RegisterUsersServiceServer(s, services.NewUsersServer(usersRepo))
	pb.RegisterOrdersServiceServer(s, services.NewOrdersServer())

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func startGatewayServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := gateway.NewGatewayMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		log.Fatalf("failed to register auth handler: %v", err)
	}
	err = pb.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		log.Fatalf("failed to register users handler: %v", err)
	}
	err = pb.RegisterOrdersServiceHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		log.Fatalf("failed to register orders handler: %v", err)
	}

	httpServer := &http.Server{
		Addr:    httpServerAddr,
		Handler: gateway.AuthenticationMiddleware(mux),
	}

	log.Printf("HTTP gateway server listening at %s", httpServerAddr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve HTTP gateway: %v", err)
	}
}

func main() {
	go startGRPCServer()

	startGatewayServer()
}
