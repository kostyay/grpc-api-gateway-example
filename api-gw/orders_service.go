package main

import (
	"context"

	pb "github.com/kostyay/grpc-api-gateway-example/api-gw/api/goclient/v1"
	usersSvcV1 "github.com/kostyay/grpc-api-gateway-example/users/api/goclient/v1"
	orderSvcV1 "github.com/kostyay/grpc-api-gateway-example/orders/api/goclient/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ordersService struct {
	usersClient  usersSvcV1.UsersServiceClient
	ordersClient orderSvcV1.OrdersServiceClient
}

func NewOrdersService(usersClient usersSvcV1.UsersServiceClient, ordersClient orderSvcV1.OrdersServiceClient) *ordersService {
	return &ordersService{
		usersClient:  usersClient,
		ordersClient: ordersClient,
	}
}

func (o *ordersService) ListOrdersWithUser(ctx context.Context, request *pb.ListOrdersWithUserRequest) (*pb.ListOrdersWithUserResponse, error) {
	// This can be done async in go routines to speed things up
	allUsers, err := o.usersClient.ListUsers(ctx, &usersSvcV1.ListUsersRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}


	allOrders, err := o.ordersClient.ListOrders(ctx, &orderSvcV1.ListOrdersRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// map users for easy mapping
	userIdToUser := map[string]*usersSvcV1.User{}
	for _, user := range allUsers.GetUsers() {
		userIdToUser[user.Id] = user
	}


	// build result
	result := &pb.ListOrdersWithUserResponse{Orders: make([]*pb.Order, len(allOrders.GetOrders()))}
	for idx, order := range allOrders.GetOrders() {
		result.Orders[idx] = &pb.Order{
			Id:        order.GetId(),
			UserId:    order.GetUserId(),
			UserEmail: userIdToUser[order.GetUserId()].GetEmail(),
			Product:   order.GetProduct(),
		}
	}

	return result, nil
}
