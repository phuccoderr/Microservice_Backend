package kafka

import (
	"cart-service/pkg/config"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func Produce(msg *sarama.ProducerMessage) error {

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewCustomPartitioner()
	cfg.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(config.Kafka.Brokers, cfg)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
		return errors.New("Failed to start Sarama producer")
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close producer:", err)
		}
	}()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
		return errors.New("Failed to send message")
	}

	log.Println(fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset))
	return nil
}
