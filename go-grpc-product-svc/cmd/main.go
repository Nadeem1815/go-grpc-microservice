package main

import (
	"fmt"
	"log"
	"net"

	// "github.com/Nadeem1815/go-grpc-api-gateway/pkg/product/pb"
	// "github.com/Nadeem1815/go-grpc-auth-svc/pkg/db"
	"github.com/Nadeem1815/go-grpc-product-svc/pkg/config"
	"github.com/Nadeem1815/go-grpc-product-svc/pkg/db"
	"github.com/Nadeem1815/go-grpc-product-svc/pkg/pb"
	"github.com/Nadeem1815/go-grpc-product-svc/pkg/services"

	// "github.com/Nadeem1815/go-grpc-product-svc/pkg/config"
	"google.golang.org/grpc"
)

// func main() {
// 	c, err := config.LoadConfig()

// 	if err != nil {
// 		log.Fatalln("Failed at config", err)
// 	}
// 	lis, err := net.Listen("tcp", c.Port)
// 	if err != nil {
// 		return

// 	}
// 	r := gin.Default()
// 	r.Run(c.Port)
// }

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}
	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
