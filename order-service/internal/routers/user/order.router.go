package user

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/middlewares"
)

type OrderRouter struct {
}

func (pr *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {

	private := Router.Group("/orders")
	private.Use(middlewares.JWTMiddleware("CUSTOMER"))
	{
		private.GET("/details")
	}
}
