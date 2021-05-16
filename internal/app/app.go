package app

import (
	"context"
	"fmt"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/00mrx00/slaves3.0_back/internal/routes"
	"github.com/00mrx00/slaves3.0_back/internal/service"
	"go.uber.org/zap"
)

func Run() {
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

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	router := routes.NewRouter(service)

	routerConfig, err := config.GetRouterConfig()

	if err != nil {
		sugar.Errorf("Failed to get Router config: %s", err.Error())
	}

	if err := router.InitRoutes().Run(routerConfig.Port); err != nil {
		fmt.Println(routerConfig)
		sugar.Errorf("Failed to initialize router: %s", err.Error())
	}

	sugar.Info("Slaves 3.0 successfully started...")
}
