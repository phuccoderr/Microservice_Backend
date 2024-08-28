package initialize

import (
	"fmt"
	"order-service/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server port", v.GetInt("server.port"))

	// config

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
