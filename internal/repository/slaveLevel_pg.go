package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type SlaveLevelPostgres struct {
	db *pgx.Conn
}

func NewSlaveLevelPostgres(db *pgx.Conn) *SlaveLevelPostgres {
	return &SlaveLevelPostgres{db: db}
}

func (rep *SlaveLevelPostgres) CreateSlaveLevel(slaveLevel domain.SlaveLevel) error {
	_, err := rep.db.Exec(context.Background(), `INSERT INTO slave_level(
		lvl,
		profit,
		money_to_update
	) VALUES ($1, $2, $3);`,
		slaveLevel.Lvl,
		slaveLevel.Profit,
		slaveLevel.MoneyToUpdate)

	if err != nil {
		return err
	}

	return nil
}

func (rep *SlaveLevelPostgres) GetSlaveLevel(lvl int32) (domain.SlaveLevel, error) {
	slaveLevel := domain.SlaveLevel{}
	err := rep.db.QueryRow(context.Background(), "SELECT * FROM slave_level WHERE lvl = $1 LIMIT 1;", lvl).Scan(
		&slaveLevel.Id,
		&slaveLevel.Lvl,
		&slaveLevel.Profit,
		&slaveLevel.MoneyToUpdate,
	)

	return slaveLevel, err
}
