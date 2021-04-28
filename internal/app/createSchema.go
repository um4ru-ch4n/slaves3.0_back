package app

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/00mrx00/slaves3.0_back/internal/domain"
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

	err = repository.CreateSchema(db)

	if err != nil {
		sugar.Errorf("Failed to create schema: %s", err.Error())
	}

	rep := repository.NewRepository(db)

	err = createUserTypes(rep)
	if err != nil {
		sugar.Errorf("Failed to create userTypes: %s", err.Error())
	}

	err = createFetters(rep)
	if err != nil {
		sugar.Errorf("Failed to create fetters: %s", err.Error())
	}

	err = createSlaveLevels(rep)
	if err != nil {
		sugar.Errorf("Failed to create slavesLevels: %s", err.Error())
	}

	err = createDefenderLevels(rep)
	if err != nil {
		sugar.Errorf("Failed to create slavesLevels: %s", err.Error())
	}

	sugar.Info("Schemas successfully created...")
}

func createUserTypes(rep *repository.Repository) error {

	err := rep.UserType.CreateUserType(domain.UserType{
		Name: "slave",
	})
	if err != nil {
		return err
	}

	err = rep.UserType.CreateUserType(domain.UserType{
		Name: "defender",
	})
	if err != nil {
		return err
	}

	err = rep.UserType.CreateUserType(domain.UserType{
		Name: "simp",
	})
	if err != nil {
		return err
	}

	err = rep.UserType.CreateUserType(domain.UserType{
		Name: "default",
	})

	return err
}

func createFetters(rep *repository.Repository) error {
	err := rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "common",
		Price:    100,
		Duration: 60 * 2,
		Cooldown: 60 * 2,
	})
	if err != nil {
		return err
	}

	err = rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "uncommon",
		Price:    10,
		Duration: 60 * 4,
		Cooldown: 60 * 4,
	})
	if err != nil {
		return err
	}

	err = rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "rare",
		Price:    12,
		Duration: 60 * 6,
		Cooldown: 60 * 6,
	})
	if err != nil {
		return err
	}

	err = rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "epic",
		Price:    14,
		Duration: 60 * 8,
		Cooldown: 60 * 8,
	})
	if err != nil {
		return err
	}

	err = rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "immortal",
		Price:    16,
		Duration: 60 * 12,
		Cooldown: 60 * 12,
	})
	if err != nil {
		return err
	}

	err = rep.Fetter.CreateFetter(domain.Fetter{
		Name:     "legendary",
		Price:    18,
		Duration: 60 * 24,
		Cooldown: 60 * 24,
	})

	return err
}

func createSlaveLevels(rep *repository.Repository) error {
	err := rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           0,
		Profit:        60,
		MoneyToUpdate: 1440,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           1,
		Profit:        120,
		MoneyToUpdate: 2880,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           2,
		Profit:        300,
		MoneyToUpdate: 7200,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           3,
		Profit:        600,
		MoneyToUpdate: 14400,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           4,
		Profit:        900,
		MoneyToUpdate: 21600,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           5,
		Profit:        1800,
		MoneyToUpdate: 43200,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           6,
		Profit:        3000,
		MoneyToUpdate: 72000,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           7,
		Profit:        6000,
		MoneyToUpdate: 144000,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           8,
		Profit:        12000,
		MoneyToUpdate: 288000,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           9,
		Profit:        30000,
		MoneyToUpdate: 720000,
	})
	if err != nil {
		return err
	}

	err = rep.SlaveLevel.CreateSlaveLevel(domain.SlaveLevel{
		Lvl:           10,
		Profit:        60000,
		MoneyToUpdate: 1440000,
	})
	if err != nil {
		return err
	}

	return err
}

func createDefenderLevels(rep *repository.Repository) error {
	err := rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            0,
		Hp:             10,
		Damage:         1,
		DamageToUpdate: 30,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            1,
		Hp:             15,
		Damage:         2,
		DamageToUpdate: 45,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            2,
		Hp:             20,
		Damage:         3,
		DamageToUpdate: 100,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            3,
		Hp:             40,
		Damage:         4,
		DamageToUpdate: 400,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            4,
		Hp:             60,
		Damage:         5,
		DamageToUpdate: 900,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            5,
		Hp:             100,
		Damage:         10,
		DamageToUpdate: 3000,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            6,
		Hp:             150,
		Damage:         115,
		DamageToUpdate: 7500,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            7,
		Hp:             200,
		Damage:         20,
		DamageToUpdate: 20000,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            8,
		Hp:             500,
		Damage:         50,
		DamageToUpdate: 100000,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            9,
		Hp:             1000,
		Damage:         100,
		DamageToUpdate: 500000,
	})
	if err != nil {
		return err
	}

	err = rep.DefenderLevel.CreateDefenderLevel(domain.DefenderLevel{
		Lvl:            10,
		Hp:             1500,
		Damage:         150,
		DamageToUpdate: 1500000,
	})
	if err != nil {
		return err
	}

	return err
}
