package initialize

import (
	"fmt"
	"mail-service/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitKafka()
	router := InitRouter()

	router.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
