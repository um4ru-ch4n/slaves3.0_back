package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type DefenderStatsPostgres struct {
	db *pgx.Conn
}

func NewDefenderStatsPostgres(db *pgx.Conn) *DefenderStatsPostgres {
	return &DefenderStatsPostgres{db: db}
}

func (rep *DefenderStatsPostgres) CreateDefenderStats(defenderStats domain.DefenderStats) (int32, error) {
	var id int32

	err := rep.db.QueryRow(context.Background(), `INSERT INTO defender_stats(
		level,
		damage_quantity
		) VALUES($1, $2)
		RETURNING id;`,
		defenderStats.Level.Id,
		defenderStats.DamageQuantity).Scan(&id)

	return id, err
}

func (rep *DefenderStatsPostgres) GetDefenderStats(id int32) (domain.DefenderStats, error) {
	defenderStats := domain.DefenderStats{Level: &domain.DefenderLevel{}}

	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			ds.id,
			ds.damage_quantity,
			dl.id,
			dl.lvl,
			dl.hp,
			dl.damage,
			dl.damage_to_update
		FROM defender_stats ds
		INNER JOIN defender_level dl
			ON dl.id = ds.level;`).Scan(
		&defenderStats.Id,
		&defenderStats.DamageQuantity,
		&defenderStats.Level.Id,
		&defenderStats.Level.Lvl,
		&defenderStats.Level.Hp,
		&defenderStats.Level.Damage,
		&defenderStats.Level.DamageToUpdate)

	return defenderStats, err
}
