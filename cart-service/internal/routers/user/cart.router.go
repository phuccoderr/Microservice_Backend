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
	cartController := controller.NewCartController(cartService, productService)

	private := router.Group("/cart")
	private.Use(middleware.JWTMiddleware("CUSTOMER"))
	{
		private.POST("/add", cartController.AddProductToCart)
		private.GET("", cartController.GetCart)
		private.DELETE("/product/:id", cartController.DeleteCart)
		private.GET("/checkout", cartController.Checkout)
		private.POST("/place_order", cartController.PlaceOrder)
	}
}
