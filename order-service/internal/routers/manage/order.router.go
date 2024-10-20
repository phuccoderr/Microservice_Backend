package manage

import (
	"github.com/gin-gonic/gin"
	"order-service/global"
	"order-service/internal/controller"
	"order-service/internal/kafka"
	"order-service/internal/middlewares"
	"order-service/internal/repository"
	"order-service/internal/service"
)

type OrderRouter struct{}

func (or *OrderRouter) InitOrderRouter(router *gin.RouterGroup) {

	orderRepository := repository.NewOrderRepo(global.Mdb)
	orderService := service.NewOrderService(orderRepository)
	redisService := service.NewOrderRedisService(global.Rdb)
	productService := service.NewProductService()
	orderController := controller.NewOrderController(orderService, redisService, productService)

	go func() {
		for {
			kafka.ConsumeOrder(redisService)
		}
	}()

	private := router.Group("/orders")
	private.Use(middlewares.JWTMiddleware("ADMIN", "SALE"))
	{
		private.GET("", orderController.GetAllOrders)
		private.GET("/:id", orderController.GetOrder)
		private.GET("/:id/status/:status", orderController.ChangeStatus)
	}
}
