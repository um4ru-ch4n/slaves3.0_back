package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type FetterService struct {
	rep repository.Fetter
}

func NewFetterService(rep repository.Fetter) *FetterService {
	return &FetterService{
		rep: rep,
	}
}

func (serv *FetterService) GetFetter(name string) (domain.Fetter, error) {
	user, err := serv.rep.GetFetter(name)

	return user, err
}

func (serv *FetterService) CreateFetter(userType domain.Fetter) error {
	err := serv.rep.CreateFetter(userType)

	return err
}
