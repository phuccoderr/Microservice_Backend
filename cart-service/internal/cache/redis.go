package cache

import (
	"cart-service/pkg/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:             config.GetRedisAddr(),
		Password:         "",
		DB:               0,
		DisableIndentity: true,
		DialTimeout:      100 * time.Millisecond,
		ReadTimeout:      100 * time.Millisecond,
		WriteTimeout:     100 * time.Millisecond,
	})

	err := Rdb.Ping(Ctx).Err()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis")
	}
}

func CartKey(customerId string) string {
	return "carts:" + customerId
}
