package user

import (
	"github.com/gin-gonic/gin"
	"order-service/global"
	"order-service/internal/controller"
	"order-service/internal/middlewares"
	"order-service/internal/repository"
	"order-service/internal/service"
)

type OrderRouter struct {
}

func (pr *OrderRouter) InitOrderRouter(router *gin.RouterGroup) {

	orderRepository := repository.NewOrderRepo(global.Mdb)
	orderService := service.NewOrderService(orderRepository)
	redisService := service.NewOrderRedisService(global.Rdb)
	productService := service.NewProductService()
	orderController := controller.NewOrderController(orderService, redisService, productService)

	private := router.Group("/orders")
	private.Use(middlewares.JWTMiddleware("CUSTOMER"))
	{
		private.GET("/details", orderController.GetOrderByCustomer)

	}
}
