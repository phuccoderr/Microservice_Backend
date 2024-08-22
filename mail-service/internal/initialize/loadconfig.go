package initialize

import (
	"github.com/spf13/viper"
	"mail-service/global"
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

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}

}
