package routes

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/cart"
	"cart-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (r Routes) CartRoute(router *gin.Engine) {

	service := cart.NewCartService()
	handler := cart.NewHandler(service)

	private := router.Group(constants.API_CART)
	private.Use(middleware.JWTMiddleware())
	{
		private.POST("/add", handler.AddProductToCart)
		private.GET("", handler.GetCart)
		private.DELETE("/product/:id", handler.DeleteCart)
		private.GET("/checkout", handler.Checkout)
		private.POST("/place_order", handler.PlaceOrder)
	}

}
