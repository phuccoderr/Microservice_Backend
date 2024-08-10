package routes

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/cart"
	"cart-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	Db *gorm.DB
}

func (r Routes) CartRoute(router *gin.Engine) {

	repository := cart.NewRepository(r.Db)
	service := cart.NewCartService(repository)
	handler := cart.NewHandler(service)

	private := router.Group(constants.API_CART)
	private.Use(middleware.JWTMiddleware())
	{
		private.POST("/add", handler.AddProductToCart)
		private.GET("", handler.GetCart)
		private.DELETE("/product/:id", handler.DeleteCart)
		private.GET("/checkout", handler.Checkout)
	}

}
