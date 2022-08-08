package services

import (
	context "context"
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

func (*UserService) AddUserVerbose(req *generators.User, stream generators.UserService_AddUserVerboseServer) error {
	stream.Send(&generators.UserResultStream{
		Status: "initial",
		User:   &generators.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&generators.UserResultStream{
		Status: "Saving",
		User: &generators.User{
			Id:    "0",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&generators.UserResultStream{
		Status: "Completed",
		User: &generators.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}
