package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joaocsv/grpc-go/grpc/generators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, failure := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if failure != nil {
		log.Fatalf("Could not connect to gRPC Server %v", failure)
	}

	defer connection.Close()

	client := generators.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client generators.UserServiceClient) {
	request := &generators.User{
		Id:    "0",
		Name:  "Joao",
		Email: "joao@joao.com",
	}

	response, failure := client.AddUser(context.Background(), request)

	if failure != nil {
		log.Fatalf("Could not make gRPC request: %v", failure)
	}

	fmt.Println(response)
}
