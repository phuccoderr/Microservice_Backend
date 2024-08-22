package initialize

import (
	"mail-service/global"
	"mail-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
