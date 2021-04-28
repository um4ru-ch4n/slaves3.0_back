package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type SlaveStatsService struct {
	rep repository.SlaveStats
}

func NewSlaveStatsService(rep repository.SlaveStats) *SlaveStatsService {
	return &SlaveStatsService{
		rep: rep,
	}
}

func (serv *SlaveStatsService) GetSlaveStats(lvl int32) (domain.SlaveStats, error) {
	user, err := serv.rep.GetSlaveStats(lvl)

	return user, err
}

func (serv *SlaveStatsService) CreateSlaveStats(slaveStats domain.SlaveStats) (int32, error) {
	id, err := serv.rep.CreateSlaveStats(slaveStats)

	return id, err
}
