package app

import (
	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/domain"
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

	defer db.Close()

	err = repository.CreateSchema(db, []interface{}{
		(*domain.UserType)(nil),
		(*domain.Fetter)(nil),
		(*domain.SlaveLevel)(nil),
		(*domain.SlaveStats)(nil),
		(*domain.DefenderLevel)(nil),
		(*domain.DefenderStats)(nil),
		(*domain.User)(nil),
		(*domain.Slaves)(nil),
	}, false)

	if err != nil {
		sugar.Errorf("Failed to create schema: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	router := routes.NewRouter(service)

	routerConfig, err := config.GetRouterConfig()

	if err != nil {
		sugar.Errorf("Failed to get Router config: %s", err.Error())
	}

	if err := router.InitRoutes().Run(routerConfig.Port); err != nil {
		sugar.Errorf("Failed to initialize router: %s", err.Error())
	}

	sugar.Info("Slaves 3.0 successfully started...")

	// err = rep.Authorization.CreateUser(domain.User{
	// 	Balance: 1000,
	// })

	// if err != nil {
	// 	sugar.Errorf("Error create user: %s", err.Error())
	// }

	// var users []domain.User
	// err = db.Model(&users).Select()
	// if err != nil {
	// 	sugar.Errorf("Error select user: %s", err.Error())
	// }

	// fmt.Println(users)
}
