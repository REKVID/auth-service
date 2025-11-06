package main

import (
	authpb "auth2/api/proto"
	"auth2/internal/auth"
	"auth2/internal/store"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Connect db
	store, err := store.NewStore("postgres://auth:auth@localhost:5432/authdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Create auth service
	svc := auth.NewService(store, "somePrivateShit")

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register our service in gRPC server
	authpb.RegisterAuthServiceServer(grpcServer, auth.NewAuthHandler(svc))

	// Start server on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error creating listener: %v", err)
	}

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
