package initialize

import (
	"cart-service/global"
	"cart-service/internal/constants"
	"cart-service/internal/middleware"
	"cart-service/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(gin.Logger())
	r.Use(middleware.CorsMiddleware())

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group(constants.API_ENPOINT)
	{
		userRouter.InitCartRouter(MainGroup)
	}
	return r
}
