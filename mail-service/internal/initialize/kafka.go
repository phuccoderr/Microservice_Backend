package initialize

import (
	"mail-service/global"
	"mail-service/internal/kafka"
)

func InitKafka() {
	client := kafka.NewKafkaClient(global.Config.Kafka)

	global.Consume = client.Reader
}
