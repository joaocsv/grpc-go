package services

import (
	context "context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/joaocsv/grpc-go/grpc/generators"
)

type UserService struct {
	generators.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, request *generators.User) (*generators.User, error) {
	return &generators.User{
		Id:    "1",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}

func (*UserService) AddUsers(stream generators.UserService_AddUsersServer) error {
	users := []*generators.User{}

	for {
		request, failure := stream.Recv()

		if failure == io.EOF {
			return stream.SendAndClose(&generators.Users{
				Users: users,
			})
		}

		if failure != nil {
			log.Fatalf("Could not receive the msg: %v", failure)
		}

		users = append(users, &generators.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		})

		fmt.Println("Adding", request.GetName())
	}
}

func (*UserService) AddUserVerbose(request *generators.User, stream generators.UserService_AddUserVerboseServer) error {
	stream.Send(&generators.UserResultStream{
		Status: "initial",
		User:   &generators.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&generators.UserResultStream{
		Status: "Saving",
		User: &generators.User{
			Id:    "0",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&generators.UserResultStream{
		Status: "Completed",
		User: &generators.User{
			Id:    "1",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUserStream(stream generators.UserService_AddUserStreamServer) error {
	for {
		request, failure := stream.Recv()

		if failure == io.EOF {
			return nil
		}

		if failure != nil {
			log.Fatalf("Error receiving stream from the client: %v", failure)
		}

		failure = stream.Send(&generators.UserResultStream{
			Status: "Added",
			User:   request,
		})

		if failure != nil {
			log.Fatalf("Error sending stream to the client: %v", failure)
		}
	}
}
