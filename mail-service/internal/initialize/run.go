package initialize

import (
	"fmt"
	"mail-service/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMailhog()
	InitKafka()
	router := InitRouter()

	router.Run(fmt.Sprintf(":%s", global.Config.Server.Port))
}
