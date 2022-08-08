package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	AddUserStream(client)
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

func AddUsers(client generators.UserServiceClient) {
	users := []*generators.User{
		{
			Id:    "1",
			Name:  "Joao1",
			Email: "Joao1@joao.com",
		},
		{
			Id:    "2",
			Name:  "Joao2",
			Email: "Joao2@joao.com",
		},
		{
			Id:    "3",
			Name:  "Joao3",
			Email: "Joao3@joao.com",
		},
		{
			Id:    "4",
			Name:  "Joao4",
			Email: "Joao4@joao.com",
		},
		{
			Id:    "5",
			Name:  "Joao5",
			Email: "Joao5@joao.com",
		},
	}

	stream, failure := client.AddUsers(context.Background())

	if failure != nil {
		log.Fatalf("Could not make gRPC request: %v", failure)
	}

	for _, user := range users {
		stream.Send(user)
		time.Sleep(time.Second * 3)
	}

	response, failure := stream.CloseAndRecv()

	if failure != nil {
		log.Fatalf("Error receiving response: %v", failure)
	}

	fmt.Println(response)
}

func AddUserVerbose(client generators.UserServiceClient) {
	request := &generators.User{
		Id:    "0",
		Name:  "Joao",
		Email: "joao@joao.com",
	}

	responseStream, failure := client.AddUserVerbose(context.Background(), request)

	if failure != nil {
		log.Fatalf("Could not make gRPC request: %v", failure)
	}

	for {
		stream, failure := responseStream.Recv()

		if failure == io.EOF {
			break
		}

		if failure != nil {
			log.Fatalf("Could not receive the msg: %v", failure)
		}

		fmt.Println("status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUserStream(client generators.UserServiceClient) {
	users := []*generators.User{
		{
			Id:    "1",
			Name:  "Joao1",
			Email: "Joao1@joao.com",
		},
		{
			Id:    "2",
			Name:  "Joao2",
			Email: "Joao2@joao.com",
		},
		{
			Id:    "3",
			Name:  "Joao3",
			Email: "Joao3@joao.com",
		},
		{
			Id:    "4",
			Name:  "Joao4",
			Email: "Joao4@joao.com",
		},
		{
			Id:    "5",
			Name:  "Joao5",
			Email: "Joao5@joao.com",
		},
	}

	stream, failure := client.AddUserStream(context.Background())

	if failure != nil {
		log.Fatalf("Could not make gRPC request: %v", failure)
	}

	wait := make(chan int)

	go func() {
		for _, user := range users {
			stream.Send(user)

			time.Sleep(time.Second * 3)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			response, failure := stream.Recv()

			if failure == io.EOF {
				break
			}

			if failure != nil {
				log.Fatalf("Error receiving data: %v", failure)
				break
			}

			fmt.Println("Receiving user", response.GetUser().GetName(), "with status:", response.GetStatus())
		}

		close(wait)
	}()

	<-wait

}
