package kafka

import (
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

func (c *Consumer) ConnectConsumer() (sarama.Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true

	fmt.Println(c.Config.Kafka.Brokers)
	return sarama.NewConsumer(c.Config.Kafka.Brokers, cfg)
}

func (c *Consumer) ConsumePartition(topic string, partition int32) {

	fmt.Println(topic)
	worker, err := c.ConnectConsumer()
	if err != nil {
		log.Fatal(err)
	}

	consumerPar, err := worker.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	fmt.Println("Consumer started!")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumerPar.Errors():
				fmt.Println(err)
			case msg := <-consumerPar.Messages():
				c.handleVerifyCustomer(msg)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Consumer stopped gracefully.")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
