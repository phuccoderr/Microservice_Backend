package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"order-service/global"
)

func LoadConfig() {
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

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
