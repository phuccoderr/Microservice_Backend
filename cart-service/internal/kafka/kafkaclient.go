package kafka

import (
	"cart-service/pkg/setting"
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func NewKafkaClient(config setting.KafkaSetting) *KafkaClient {
	return &KafkaClient{
		Writer: getKafkaWriter(config.Brokers, config.Topic),
	}
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
