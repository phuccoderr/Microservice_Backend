package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	DB    DBConfig
	Kafka KafkaConfig
	Redis RedisConfig
	Jwt   JwtConfig
)

type JwtConfig struct {
	Secret      string
	ExpiredTime int
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type KafkaConfig struct {
	Brokers []string
}

type RedisConfig struct {
	Host string
	Port string
}

func LoadConfig() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	DB = DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	Kafka = KafkaConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKERS")},
	}
	Redis = RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	expired, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		log.Fatal("Error parsing JWT_EXPIRATION")
	}
	Jwt = JwtConfig{
		Secret:      os.Getenv("JWT_SECRET"),
		ExpiredTime: expired,
	}
}

func GetPostgresURL() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DB.Host, DB.User, DB.Password, DB.DBName, DB.Port, DB.SSLMode)

	return dsn
}

func GetRedisAddr() string {
	dsn := fmt.Sprintf("%s:%s", Redis.Host, Redis.Port)
	return dsn
}
