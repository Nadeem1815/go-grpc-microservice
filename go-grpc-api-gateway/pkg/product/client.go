package product

import (
	"fmt"

	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/config"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

// we need to define the client to communicate with the Product Microservice.
type ServiceClient struct {
	client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("could not connect:", err)

	}
	return pb.NewProductServiceClient(cc)
}
