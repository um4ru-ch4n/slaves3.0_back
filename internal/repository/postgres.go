package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewPostgresDB(cfg config.DbConfig) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Addr,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateSchema(db *pg.DB, models []interface{}, isTemp bool) error {
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        isTemp,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
