package app

import (
	"fmt"
	"os"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/logger"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/00mrx00/slaves3.0_back/internal/routes"
	"github.com/00mrx00/slaves3.0_back/internal/service"
	"go.uber.org/zap"
)

func Run() {
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

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	router := routes.NewRouter(service)

	routerConfig, err := config.GetRouterConfig()

	if err != nil {
		logger.Fatal("Failed to get Router config: ", zap.Error(err))
	}

	if err := router.InitRoutes().Run(routerConfig.Port); err != nil {
		logger.Fatal("Failed to initialize router: ", zap.Error(err))
	}

	logger.Info("Slaves 3.0 successfully started...")
}
