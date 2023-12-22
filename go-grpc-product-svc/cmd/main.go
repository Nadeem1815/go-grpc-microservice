package main

import (
	"log"

	"github.com/Nadeem1815/go-grpc-product-svc/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// lis,err:=net.Listen("tcp",c.Port)
	// if err!=nil {
	// 	return

	// }
	// r := gin.Default()
	// r.Run(c.Port)
}
