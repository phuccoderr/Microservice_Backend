package main

import (
	"mail-service/config"
	"mail-service/internal/kafka"
)

const webPort = "9160"

func main() {
	config := config.LoadConfig()

	consumer := kafka.Consumer{
		Config: config,
	}
	consumer.ConsumePartition("customer-verify-events-topics", 0)
}
