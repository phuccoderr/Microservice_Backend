package initialize

import (
	"github.com/gin-gonic/gin"
	"mail-service/internal/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	router.NewRouter(r)
	return r
}
