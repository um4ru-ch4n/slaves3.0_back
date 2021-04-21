package repository

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/go-pg/pg/v10"
)

type Authorization interface {
	CreateUser(user domain.User) error
	GetUser(id int) (domain.User, error)
}

type Repository struct {
	Authorization Authorization
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
