package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type UserTypeService struct {
	rep repository.UserType
}

func NewUserTypeService(rep repository.UserType) *UserTypeService {
	return &UserTypeService{
		rep: rep,
	}
}

func (serv *UserTypeService) GetUserType(name string) (domain.UserType, error) {
	user, err := serv.rep.GetUserType(name)

	return user, err
}

func (serv *UserTypeService) CreateUserType(userType domain.UserType) error {
	err := serv.rep.CreateUserType(userType)

	return err
}
