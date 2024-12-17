package main

import (
	userv1 "go-server/gen/user/v1"
	"go-server/internal"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50551")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	userv1.RegisterUserServiceServer(server, internal.NewUserService())

	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
