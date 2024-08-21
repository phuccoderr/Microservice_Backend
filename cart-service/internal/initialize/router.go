package initialize

import (
	"cart-service/global"
	"cart-service/internal/constants"
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

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group(constants.API_ENPOINT)
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitCartRouter(MainGroup)
	}
	return r
}
