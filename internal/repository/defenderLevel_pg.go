package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type DefenderLevelPostgres struct {
	db *pgx.Conn
}

func NewDefenderLevelPostgres(db *pgx.Conn) *DefenderLevelPostgres {
	return &DefenderLevelPostgres{db: db}
}

func (rep *DefenderLevelPostgres) CreateDefenderLevel(defenderLevel domain.DefenderLevel) error {
	_, err := rep.db.Exec(context.Background(), `INSERT INTO defender_level(
		lvl,
		hp,
		damage,
		damage_to_update
	) VALUES ($1, $2, $3, $4);`,
		defenderLevel.Lvl,
		defenderLevel.Hp,
		defenderLevel.Damage,
		defenderLevel.DamageToUpdate)

	return err
}

func (rep *DefenderLevelPostgres) GetDefenderLevel(lvl int32) (domain.DefenderLevel, error) {
	defenderLevel := domain.DefenderLevel{}
	err := rep.db.QueryRow(context.Background(), "SELECT * FROM defender_level WHERE lvl=$1 LIMIT 1;", lvl).Scan(
		&defenderLevel.Id,
		&defenderLevel.Lvl,
		&defenderLevel.Hp,
		&defenderLevel.Damage,
		&defenderLevel.DamageToUpdate)

	return defenderLevel, err
}
