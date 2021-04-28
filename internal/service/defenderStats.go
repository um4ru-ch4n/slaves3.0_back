package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type DefenderStatsService struct {
	rep repository.DefenderStats
}

func NewDefenderStatsService(rep repository.DefenderStats) *DefenderStatsService {
	return &DefenderStatsService{
		rep: rep,
	}
}

func (serv *DefenderStatsService) GetDefenderStats(lvl int32) (domain.DefenderStats, error) {
	defStats, err := serv.rep.GetDefenderStats(lvl)

	return defStats, err
}

func (serv *DefenderStatsService) CreateDefenderStats(defenderStats domain.DefenderStats) (int32, error) {
	id, err := serv.rep.CreateDefenderStats(defenderStats)

	return id, err
}
