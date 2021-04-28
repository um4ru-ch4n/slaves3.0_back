package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type DefenderLevelService struct {
	rep repository.DefenderLevel
}

func NewDefenderLevelService(rep repository.DefenderLevel) *DefenderLevelService {
	return &DefenderLevelService{
		rep: rep,
	}
}

func (serv *DefenderLevelService) GetDefenderLevel(lvl int32) (domain.DefenderLevel, error) {
	user, err := serv.rep.GetDefenderLevel(lvl)

	return user, err
}

func (serv *DefenderLevelService) CreateDefenderLevel(defenderLevel domain.DefenderLevel) error {
	err := serv.rep.CreateDefenderLevel(defenderLevel)

	return err
}
