package config

import (
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
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
}

func LoadConfig() *Config {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//} use Docker , dont need godotenv
	return &Config{
		Mailer: CreateMail(),
		Kafka: KafkaConfig{
			Brokers: []string{os.Getenv("KAFKA_BROKERS")},
			Topics: []string{os.Getenv("KAFKA_TOPIC_VERIFY_ACCOUNT"),
				os.Getenv("KAFKA_TOPIC_VERIFY_PASSWORD"),
				os.Getenv("KAFKA_TOPIC_PLACE_ORDER")},
		},
	}
}

func CreateMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:     os.Getenv("MAIL_DOMAIN"),
		Host:       os.Getenv("MAIL_HOST"),
		Port:       port,
		Encryption: os.Getenv("MAIL_ENCRYPTION"),
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
	}

	return m
}
