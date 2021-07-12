package config

import "github.com/spf13/viper"

var config *viper.Viper

func Init() *viper.Viper {
	config = viper.New()
	config.SetConfigType("json")
	config.SetConfigName("config")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}

func Get() *viper.Viper {
	return config
}
