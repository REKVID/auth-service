package main

import (
	"context"
	"log"
	"net/http"

	authpb "auth2/api/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//  HTTP роутер
	mux := runtime.NewServeMux()

	// Connect to gRPC 
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Register HTTP routes from proto
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051",
		opts,
	)
	if err != nil {
		log.Fatalf("Error registering gateway: %v", err)
	}

	log.Println("HTTP Gateway running on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}