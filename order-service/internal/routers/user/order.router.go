package user

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/kafka"
)

type OrderRouter struct {
}

func (pr *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	for {
		kafka.ConsumeOrder()
	}
	//public
	routerPublic := Router.Group("/order")
	{
		routerPublic.GET("")
	}
	//private
}
