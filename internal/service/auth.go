package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{
		rep: rep,
	}
}

func (serv *AuthService) CreateUser(user domain.User) error {
	return nil
}

func (serv *AuthService) GetUser(id int) (domain.User, error) {
	return domain.User{}, nil
}
