package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
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
		return RouterConfig{}, errors.Wrap(err, "failed to read yml router config")
	}

	if err := godotenv.Load(); err != nil {
		return RouterConfig{}, errors.Wrap(err, "failed to load env router")
	}

	return RouterConfig{
		Port: os.Getenv("PORT"),
	}, nil

}
