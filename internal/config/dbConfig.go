package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func GetDbConfig() (DbConfig, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		return DbConfig{}, err
	}

	if err := godotenv.Load(); err != nil {
		return DbConfig{}, err
	}

	return DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.ssl"),
	}, nil

}
