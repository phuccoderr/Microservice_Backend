package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
	Kafka  KafkaConfig
}

type KafkaConfig struct {
	Brokers []string
	Topics  []string
}
type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		Mailer: CreateMail(),
		Kafka: KafkaConfig{
			Brokers: []string{"localhost:9092"},
			Topics: []string{os.Getenv("KAFKA_TOPIC_VERIFY_ACCOUNT"),
				os.Getenv("KAFKA_TOPIC_VERIFY_PASSWORD")},
		},
	}
}

func CreateMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return m
}
