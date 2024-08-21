package initialize

import (
	"cart-service/global"
	"fmt"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitRedis()
	InitKafka()
	router := InitRouter()

	router.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
