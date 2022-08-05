package services

import (
	context "context"

	"github.com/joaocsv/grpc-go/grpc/generators"
)

type UserService struct {
	generators.UnimplementedUserServiceServer
}

func AddUser(context context.Context, user *generators.User) (*generators.User, error) {
	return &generators.User{
		Id:    "1",
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil
}
