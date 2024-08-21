package user

import (
	"cart-service/global"
	"cart-service/internal/controller"
	"cart-service/internal/middleware"
	"cart-service/internal/service"
	"github.com/gin-gonic/gin"
)

type CartRouter struct {
}

func (cr *CartRouter) InitCartRouter(router *gin.RouterGroup) {
	productService := service.NewProductService()
	cartService := service.NewCartService(global.Rdb)
	controller := controller.NewCartController(cartService, productService)

	private := router.Group("/cart")
	private.Use(middleware.JWTMiddleware())
	{
		private.POST("/add", controller.AddProductToCart)
		private.GET("", controller.GetCart)
		private.DELETE("/product/:id", controller.DeleteCart)
		private.GET("/checkout", controller.Checkout)
		private.POST("/place_order", controller.PlaceOrder)
	}
}
