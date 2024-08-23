package initialize

import (
	"order-service/global"
	"order-service/internal/cache"
)

func InitRedis() {
	client := cache.NewRedisClient(global.Config.Redis)
	global.Rdb = client
}
