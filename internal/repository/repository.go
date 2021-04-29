package repository

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type Authorization interface {
	CreateUser(user domain.User) error
	GetUser(id int32) (domain.User, error)
	GetUserType(userId int32) (string, error)
	GetFriendInfoLocal(id int32) (domain.FriendInfoLocal, error)
}

type UserType interface {
	CreateUserType(userType domain.UserType) error
	GetUserType(name string) (domain.UserType, error)
}

type Fetter interface {
	CreateFetter(fetter domain.Fetter) error
	GetFetter(name string) (domain.Fetter, error)
}

type SlaveLevel interface {
	CreateSlaveLevel(slaveLevel domain.SlaveLevel) error
	GetSlaveLevel(lvl int32) (domain.SlaveLevel, error)
}

type SlaveStats interface {
	CreateSlaveStats(slaveStats domain.SlaveStats) (int32, error)
	GetSlaveStats(id int32) (domain.SlaveStats, error)
}

type DefenderLevel interface {
	CreateDefenderLevel(defenderLevel domain.DefenderLevel) error
	GetDefenderLevel(lvl int32) (domain.DefenderLevel, error)
}

type DefenderStats interface {
	CreateDefenderStats(defenderStats domain.DefenderStats) (int32, error)
	GetDefenderStats(id int32) (domain.DefenderStats, error)
}

type Repository struct {
	Authorization Authorization
	UserType      UserType
	Fetter        Fetter
	SlaveLevel    SlaveLevel
	SlaveStats    SlaveStats
	DefenderLevel DefenderLevel
	DefenderStats DefenderStats
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserType:      NewUserTypePostgres(db),
		Fetter:        NewFetterPostgres(db),
		SlaveLevel:    NewSlaveLevelPostgres(db),
		SlaveStats:    NewSlaveStatsPostgres(db),
		DefenderLevel: NewDefenderLevelPostgres(db),
		DefenderStats: NewDefenderStatsPostgres(db),
	}
}
