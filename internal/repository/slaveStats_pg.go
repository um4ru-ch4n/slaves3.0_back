package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type SlaveStatsPostgres struct {
	db *pgx.Conn
}

func NewSlaveStatsPostgres(db *pgx.Conn) *SlaveStatsPostgres {
	return &SlaveStatsPostgres{db: db}
}

func (rep *SlaveStatsPostgres) CreateSlaveStats(slaveStats domain.SlaveStats) (int32, error) {
	var id int32

	err := rep.db.QueryRow(context.Background(), `INSERT INTO slave_stats(
		level,
		money_quantity
	) VALUES($1, $2)
		RETURNING id;`,
		slaveStats.Level.Id,
		slaveStats.MoneyQuantity).Scan(&id)

	return id, err
}

func (rep *SlaveStatsPostgres) GetSlaveStats(id int32) (domain.SlaveStats, error) {
	slaveStats := domain.SlaveStats{Level: &domain.SlaveLevel{}}

	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			ss.id,
			ss.money_quantity,
			sl.id,
			sl.lvl,
			sl.profit,
			sl.money_to_update
		FROM slave_stats ss
		INNER JOIN slave_level sl
			ON sl.id = ss.level LIMIT 1;`).Scan(
		&slaveStats.Id,
		&slaveStats.MoneyQuantity,
		&slaveStats.Level.Id,
		&slaveStats.Level.Lvl,
		&slaveStats.Level.Profit,
		&slaveStats.Level.MoneyToUpdate)

	return slaveStats, err
}
