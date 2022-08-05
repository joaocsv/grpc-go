package main

import (
	"log"
	"net"

	"github.com/joaocsv/grpc-go/grpc/generators"
	"github.com/joaocsv/grpc-go/grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, failure := net.Listen("tcp", "localhost:50051")

	if failure != nil {
		log.Fatalf("Could not connect: %v", failure)
	}

	grpcServer := grpc.NewServer()

	generators.RegisterUserServiceServer(grpcServer, &services.UserService{})
	reflection.Register(grpcServer)

	failure = grpcServer.Serve(listener)

	if failure != nil {
		log.Fatalf("Could not serve: %v", failure)
	}
}
