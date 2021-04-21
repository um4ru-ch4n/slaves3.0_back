package repository

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/go-pg/pg/v10"
)

type AuthPostgres struct {
	db *pg.DB
}

func NewAuthPostgres(db *pg.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (rep *AuthPostgres) CreateUser(user domain.User) error {
	_, err := rep.db.Model(&user).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (rep *AuthPostgres) GetUser(id int) (domain.User, error) {
	user := &domain.User{}

	err := rep.db.Model(user).Where("id = ?", id).Select()

	return *user, err
}
