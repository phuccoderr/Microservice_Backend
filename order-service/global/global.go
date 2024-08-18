package global

import (
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"order-service/pkg/logger"
	"order-service/pkg/setting"
)

var (
	Config   setting.Config
	Logger   *logger.LoggerZap
	Mdb      *gorm.DB
	Producer *kafka.Writer
)
