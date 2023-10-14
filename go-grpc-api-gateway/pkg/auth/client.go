package auth

import (
	"fmt"

	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)

	}
	return pb.NewAuthServiceClient(cc)
}
