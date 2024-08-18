package initialize

import (
	"github.com/gin-gonic/gin"
	"order-service/global"
	"order-service/internal/routers"
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

	r.Use() //logging
	r.Use() //cross
	r.Use() //limiter global

	//manage := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("api/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitOrderRouter(MainGroup)
	}
	return r
}
