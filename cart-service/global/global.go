package global

import (
	"cart-service/pkg/logger"
	"cart-service/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config  setting.Config
	Logger  *logger.LoggerZap
	Rdb     *redis.Client
	Produce *kafka.Writer
)
