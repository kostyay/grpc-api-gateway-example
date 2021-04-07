package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/kostyay/grpc-api-gateway-example/api-gw/api/goclient/v1"
	ordersSvcV1 "github.com/kostyay/grpc-api-gateway-example/orders/api/goclient/v1"
	usersSvcV1 "github.com/kostyay/grpc-api-gateway-example/users/api/goclient/v1"
	"google.golang.org/grpc"
)

const (
	listenAddress = "0.0.0.0:9090"
	ordersSvc = "orders:9090"
	usersSvc = "users:9090"
)

func newOrdersSvcClient() (ordersSvcV1.OrdersServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), ordersSvc, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("orders client: %w", err)
	}

	return ordersSvcV1.NewOrdersServiceClient(conn), nil
}

func newUsersSvcClient() (usersSvcV1.UsersServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), usersSvc, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("users client: %w", err)
	}

	return usersSvcV1.NewUsersServiceClient(conn), nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	log.Printf("APIGW service starting on %s", listenAddress)

	// connect to orders svc
	ordersClient, err := newOrdersSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to users svc
	usersClient, err := newUsersSvcClient()
	if err != nil {
		panic(err)
	}


	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	pb.RegisterOrdersServiceServer(s, NewOrdersService(usersClient, ordersClient))
	pb.RegisterUsersServiceServer(s, NewUsersService(usersClient))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}