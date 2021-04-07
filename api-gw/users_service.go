package main

import (
	"context"

	pb "github.com/kostyay/grpc-api-gateway-example/api-gw/api/goclient/v1"
	usersSvcV1 "github.com/kostyay/grpc-api-gateway-example/users/api/goclient/v1"
)

type usersService struct {
	usersClient usersSvcV1.UsersServiceClient
}

func NewUsersService(usersClient usersSvcV1.UsersServiceClient) *usersService {
	return &usersService{
		usersClient: usersClient,
	}
}

func (u *usersService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := u.usersClient.CreateUser(ctx, &usersSvcV1.CreateUserRequest{Email: request.GetEmail()})
	if err != nil {
		return nil, err
	}

	// as you can see in this case the messages are quite similar (for now) but we have to translate
	// them between API structs and internal structs
	return &pb.CreateUserResponse{User: &pb.User{
		Id:    res.GetUser().GetId(),
		Email: res.GetUser().GetEmail(),
	}}, nil
}

func (u *usersService) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	panic("implement me")
}


