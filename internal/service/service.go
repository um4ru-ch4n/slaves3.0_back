package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) error
	GetUser(id int) (domain.User, error)
}

type Service struct {
	Authorization Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
