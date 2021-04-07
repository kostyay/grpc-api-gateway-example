package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kostyay/grpc-api-gateway-example/orders/api/goclient/v1"
	"google.golang.org/grpc"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type ordersService struct {}

func (o *ordersService) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return &pb.ListOrdersResponse{Orders: []*pb.Order{
		{
			Id:      "o1",
			UserId:  "1",
			Product: "product-1",
		},
		{
			Id:      "o2",
			UserId:  "1",
			Product: "product-2",
		},
	}}, nil
}

func main() {
	log.Printf("Orders service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrdersServiceServer(s, &ordersService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}