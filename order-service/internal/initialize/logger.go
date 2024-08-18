package initialize

import (
	"order-service/global"
	"order-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
