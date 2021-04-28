package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type SlaveLevelService struct {
	rep repository.SlaveLevel
}

func NewSlaveLevelService(rep repository.SlaveLevel) *SlaveLevelService {
	return &SlaveLevelService{
		rep: rep,
	}
}

func (serv *SlaveLevelService) GetSlaveLevel(lvl int32) (domain.SlaveLevel, error) {
	user, err := serv.rep.GetSlaveLevel(lvl)

	return user, err
}

func (serv *SlaveLevelService) CreateSlaveLevel(slaveLevel domain.SlaveLevel) error {
	err := serv.rep.CreateSlaveLevel(slaveLevel)

	return err
}
