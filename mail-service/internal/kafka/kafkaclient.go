package kafka

import (
	"github.com/segmentio/kafka-go"
	"mail-service/pkg/setting"
	"strings"
	"time"
)

type KafkaClient struct {
	Reader *kafka.Reader
}

func NewKafkaClient(config setting.KafkaSetting) *KafkaClient {
	return &KafkaClient{
		Reader: getKafkaReader(config.Brokers, config.Topics, config.Group),
	}
}

func getKafkaReader(kafkaURL string, topics []string, group string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        group,
		GroupTopics:    topics,
		MinBytes:       10e3,
		MaxBytes:       10e3,
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}
