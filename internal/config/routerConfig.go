package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type RouterConfig struct {
	Port string
}

func GetRouterConfig() (RouterConfig, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		return RouterConfig{}, err
	}

	if err := godotenv.Load(); err != nil {
		return RouterConfig{}, err
	}

	return RouterConfig{
		Port: viper.GetString("router.port"),
	}, nil

}
