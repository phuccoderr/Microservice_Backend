package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"order-service/pkg/logger"
	"order-service/pkg/setting"
)

var (
	Config   setting.Config
	Logger   *logger.LoggerZap
	Mdb      *gorm.DB
	Rdb      *redis.Client
	Producer *kafka.Writer
	Consumer *kafka.Reader
)
