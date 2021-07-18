package config

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() *viper.Viper {
	config = viper.New()
	config.SetConfigType("json")
	config.SetConfigName("config")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return config
}

func Get() *viper.Viper {
	return config
}

func GetJWTAlgorithm() jwt.SigningMethod {
	algo := config.GetString("jwt.algo")
	switch algo {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "RS256":
		return jwt.SigningMethodRS256
	default:
		return jwt.SigningMethodHS256
	}
}
