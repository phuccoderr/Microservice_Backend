package routers

import (
	"github.com/gin-gonic/gin"
	c "order-service/internal/controller"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Group("api/v1/order")
	{
		r.GET("", c.NewOrderController().GetOrderList)
	}

	return r
}
