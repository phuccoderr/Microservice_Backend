package initialize

import (
	"cart-service/global"
	"cart-service/internal/kafka"
)

func InitKafka() {
	client := kafka.NewKafkaClient(global.Config.Kafka)
	global.Produce = client.Writer
}
