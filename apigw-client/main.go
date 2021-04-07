package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	apiClient "github.com/kostyay/grpc-api-gateway-example/api-gw/api/goclient/v1"
	"google.golang.org/grpc"
)

const apiSvc = "localhost:9090"

func main() {
	conn, err := grpc.DialContext(context.TODO(), apiSvc, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	api := apiClient.NewOrdersServiceClient(conn)

	res, err := api.ListOrdersWithUser(context.Background(), &apiClient.ListOrdersWithUserRequest{})
	if err != nil {
		panic(err)
	}

	resp, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		panic(err)
	}

	fmt.Printf(resp)
}
