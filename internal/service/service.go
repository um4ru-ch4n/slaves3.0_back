package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type Authorization interface {
	GetUser(id int32) (domain.User, error)
	CreateUser(user domain.User) error
	GetUserVkInfo(token string) (domain.UserVkInfo, error)
	GetFriendsList(token string, friendId int32) ([]domain.FriendInfo, error)
	GetFriendInfoLocal(friendId int32) (domain.FriendInfoLocal, error)
}

type UserType interface {
	GetUserType(name string) (domain.UserType, error)
	CreateUserType(userType domain.UserType) error
}

type SlaveLevel interface {
	GetSlaveLevel(lvl int32) (domain.SlaveLevel, error)
	CreateSlaveLevel(slaveLevel domain.SlaveLevel) error
}

type SlaveStats interface {
	GetSlaveStats(id int32) (domain.SlaveStats, error)
	CreateSlaveStats(slaveStats domain.SlaveStats) (int32, error)
}

type DefenderLevel interface {
	GetDefenderLevel(lvl int32) (domain.DefenderLevel, error)
	CreateDefenderLevel(defenderLevel domain.DefenderLevel) error
}

type DefenderStats interface {
	GetDefenderStats(id int32) (domain.DefenderStats, error)
	CreateDefenderStats(defenderStats domain.DefenderStats) (int32, error)
}

type Fetter interface {
	GetFetter(name string) (domain.Fetter, error)
	CreateFetter(fetter domain.Fetter) error
}

type Service struct {
	Authorization Authorization
	UserType      UserType
	SlaveLevel    SlaveLevel
	SlaveStats    SlaveStats
	DefenderLevel DefenderLevel
	DefenderStats DefenderStats
	Fetter        Fetter
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		UserType:      NewUserTypeService(rep.UserType),
		SlaveLevel:    NewSlaveLevelService(rep.SlaveLevel),
		SlaveStats:    NewSlaveStatsService(rep.SlaveStats),
		DefenderLevel: NewDefenderLevelService(rep.DefenderLevel),
		DefenderStats: NewDefenderStatsService(rep.DefenderStats),
		Fetter:        NewFetterService(rep.Fetter),
	}
}
