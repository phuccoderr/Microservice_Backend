package kafka

import (
	"github.com/segmentio/kafka-go"
	"order-service/pkg/setting"
	"strings"
	"time"
)

type KafkaClient struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func NewKafkaClient(config setting.KafkaSetting) *KafkaClient {
	return &KafkaClient{
		Reader: getKafkaReader(config.Brokers, config.Topic, "go-groupA"),
		Writer: getKafkaWriter(config.Brokers, config.Topic),
	}
}

func getKafkaReader(kafkaURL string, topic string, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,
		MaxBytes:       10e3,
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
