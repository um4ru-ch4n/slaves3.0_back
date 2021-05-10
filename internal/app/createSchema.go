package app

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"go.uber.org/zap"
)

func CreateSchema() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	dbConfig, err := config.GetDbConfig()

	if err != nil {
		sugar.Errorf("Failed to get DB config db: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(dbConfig)

	if err != nil {
		sugar.Errorf("Failed to connect DB: %s", err.Error())
	}

	defer db.Close(context.Background())

	if err := repository.CreateSchema(db); err != nil {
		sugar.Errorf("Failed to create schema: %s", err.Error())
	}

	if err := repository.CreateUserTypes(db); err != nil {
		sugar.Errorf("Failed to create user_type: %s", err.Error())
	}

	if err := repository.CreateFetter(db); err != nil {
		sugar.Errorf("Failed to create fetter: %s", err.Error())
	}

	sugar.Info("Schemas successfully created...")
}
