package config

import (
	"os"

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

	// if err := godotenv.Load(); err != nil {
	// 	return DbConfig{}, err
	// }

	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.ssl"),
	}, nil

}
