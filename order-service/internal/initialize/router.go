package initialize

import (
	"github.com/gin-gonic/gin"
	"order-service/global"
	"order-service/internal/constants"
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

	r.Use(gin.Logger()) //logging

	manage := routers.RouterGroupApp.Manage

	MainGroup := r.Group(constants.API_ENPOINT)
	{
		MainGroup.GET("/checkStatus")
	}
	{
		manage.InitOrderRouter(MainGroup)
	}
	return r
}
