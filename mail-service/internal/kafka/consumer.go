package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"mail-service/config"
	"os"
	"os/signal"
	"syscall"
)

type Consumer struct {
	Config *config.Config
}

func (c *Consumer) ConsumeKafka() {
	cfgKafka := sarama.NewConfig()
	cfgKafka.Consumer.Return.Errors = true
	cfgKafka.Consumer.Offsets.Initial = sarama.OffsetOldest

	fmt.Println(c.Config.Kafka)

	consumer, err := sarama.NewConsumerGroup(c.Config.Kafka.Brokers, "my-group", cfgKafka)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}

	handler := &ConsumerGroupHandler{
		Config: c.Config,
	}

	if handler.Config == nil {
		log.Fatalf("ConsumerGroupHandler config is nil")
	}

	go func() {
		for {
			ctx := context.Background()
			if err := consumer.Consume(ctx, c.Config.Kafka.Topics, handler); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	if err := consumer.Close(); err != nil {
		log.Fatalf("Error closing consumer: %v", err)
	}

}
