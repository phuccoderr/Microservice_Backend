package initialize

import (
	"order-service/global"
	"order-service/internal/kafka"
)

func IniteKafka() {

	client := kafka.NewKafkaClient(global.Config.Kafka)
	global.Producer = client.Writer
	global.Consumer = client.Reader

}
