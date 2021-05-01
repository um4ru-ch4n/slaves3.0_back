package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type UserTypePostgres struct {
	db *pgx.Conn
}

func NewUserTypePostgres(db *pgx.Conn) *UserTypePostgres {
	return &UserTypePostgres{db: db}
}

func (rep *UserTypePostgres) CreateUserType(userType domain.UserType) error {
	_, err := rep.db.Exec(context.Background(), "INSERT INTO user_type(name) VALUES ($1);", userType.Name)

	return err
}

func (rep *UserTypePostgres) GetUserType(name string) (domain.UserType, error) {
	userType := domain.UserType{}
	err := rep.db.QueryRow(context.Background(), "SELECT * FROM user_type WHERE name = $1 LIMIT 1;", name).Scan(&userType.Id, &userType.Name)

	return userType, err
}
