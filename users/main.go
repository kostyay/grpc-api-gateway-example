package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kostyay/grpc-api-gateway-example/users/api/goclient/v1"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type usersService struct {

}

func (u *usersService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{User: &pb.User{
		Id:    "1",
		Email: "user1@email.com",
	}}, nil
}

func (u *usersService) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return &pb.ListUsersResponse{
		Users: []*pb.User{
			{
				Id:    "1",
				Email: "user1@email.com",
			},
			{
				Id:    "2",
				Email: "user2@email.com",
			},
		},
	}, nil
}

func main() {
	log.Printf("Users service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &usersService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}