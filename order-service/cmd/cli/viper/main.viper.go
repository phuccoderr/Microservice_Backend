package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server port", v.GetInt("server.port"))

	// config
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	fmt.Println("Server port", config.Server.Port)
}
