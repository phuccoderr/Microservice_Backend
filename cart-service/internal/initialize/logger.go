package initialize

import (
	"cart-service/global"
	"cart-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
