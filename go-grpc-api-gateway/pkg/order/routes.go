package order

import (
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/auth"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/config"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	// svc := &ServiceClient{
	// 	client: InitServiceClient(c),
	// }
	svc := &ServiceClient{
		client: InitServiceClient(c),
	}
	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.client)
}
