package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type FetterPostgres struct {
	db *pgx.Conn
}

func NewFetterPostgres(db *pgx.Conn) *FetterPostgres {
	return &FetterPostgres{db: db}
}

func (rep *FetterPostgres) CreateFetter(fetter domain.Fetter) error {
	_, err := rep.db.Exec(context.Background(), `INSERT INTO fetter(
		name,
		price,
		duration,
		cooldown
	) VALUES ($1, $2, $3, $4);`,
		fetter.Name,
		fetter.Price,
		fetter.Duration,
		fetter.Cooldown)

	return err
}

func (rep *FetterPostgres) GetFetter(name string) (domain.Fetter, error) {
	fetter := domain.Fetter{}
	err := rep.db.QueryRow(context.Background(), "SELECT * FROM fetter WHERE name = $1", name).Scan(
		&fetter.Id,
		&fetter.Name,
		&fetter.Price,
		&fetter.Duration,
		&fetter.Cooldown)

	return fetter, err
}
