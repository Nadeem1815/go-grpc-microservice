package order

import (
	"fmt"

	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/config"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

// the order microservice need to have a client as well

type ServiceClient struct {
	client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// using WithInsecure() because no SSl running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("could not connect:", err)

	}
	return pb.NewOrderServiceClient(cc)
}
