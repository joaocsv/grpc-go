package main

import (
	"log"
	"net"

	"github.com/joaocsv/grpc-go/grpc/generators"
	"github.com/joaocsv/grpc-go/grpc/services"
	"google.golang.org/grpc"
)

func main() {
	listener, failure := net.Listen("tcp", "0.0.0.0:5555")

	if failure != nil {
		log.Fatalf("Could not connect: %v", failure)
	}

	grpcServer := grpc.NewServer()

	generators.RegisterUserServiceServer(grpcServer, services.UserService{})

	failure = grpcServer.Serve(listener)

	if failure != nil {
		log.Fatalf("Could not serve: %v", failure)
	}
}
