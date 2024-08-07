package main

import (
	"mail-service/config"
	"mail-service/internal/kafka"
)

func main() {
	config := config.LoadConfig()

	consumer := kafka.Consumer{
		Config: config,
	}
	consumer.ConsumeKafka()
}
