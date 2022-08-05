package services

import (
	context "context"

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
