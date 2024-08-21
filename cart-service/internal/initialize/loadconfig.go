package initialize

import (
	"cart-service/global"
	"fmt"
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
