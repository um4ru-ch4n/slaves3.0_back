package app

import (
	"fmt"
	"os"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/logger"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"go.uber.org/zap"
)

func CreateSchema() {
	logger, err := logger.NewLogger()
	if err != nil {
		fmt.Println("logger create failed: ", err)
		os.Exit(1)
	}

	dbConfig, err := config.GetDbConfig()

	if err != nil {
		logger.Fatal("Failed to get DB config db: ", zap.Error(err))
	}

	db, err := repository.NewPostgresDB(dbConfig)

	if err != nil {
		logger.Fatal("Failed to connect DB: ", zap.Error(err))
	}

	defer db.Close()

	if err := repository.CreateSchema(db); err != nil {
		logger.Fatal("Failed to create schema: ", zap.Error(err))
	}

	if err := repository.CreateUserTypes(db); err != nil {
		logger.Fatal("Failed to create user_type: ", zap.Error(err))
	}

	if err := repository.CreateFetter(db); err != nil {
		logger.Fatal("Failed to create fetter: ", zap.Error(err))
	}

	logger.Info("Schemas successfully created...")
}
