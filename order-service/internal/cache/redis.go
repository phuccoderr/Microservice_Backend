package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"order-service/global"
	"order-service/pkg/setting"
	"time"
)

func NewRedisClient(config setting.RedisSetting) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:             fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password:         "",
		DB:               0,
		DisableIndentity: true,
		DialTimeout:      100 * time.Millisecond,
		ReadTimeout:      100 * time.Millisecond,
		WriteTimeout:     100 * time.Millisecond,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		global.Logger.Info("Connected to Redis")
	}

	return redisClient
}

func CartKey(customerId string) string {
	return "carts:" + customerId
}
