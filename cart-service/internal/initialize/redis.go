package initialize

import (
	"cart-service/global"
	"cart-service/internal/cache"
)

func InitRedis() {
	global.Rdb = cache.NewRedisClient(global.Config.Redis)
}
