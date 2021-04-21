package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Addr     string
	User     string
	Password string
	Database string
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
		Addr:     viper.GetString("db.port"),
		User:     os.Getenv("DB_USERNAME"),
		Database: viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	}, nil

}
