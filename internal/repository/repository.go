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
	SlaveBuyUpdateInfo(newData domain.SlaveBuyUpdateInfo) error
	SlaveCountBalanceUpdate(userId int32, slavesCount int32, balance int64, gold int32) error
	SetHasFetter(userId int32, hasFetter bool) error
	GetUserBalance(userId int32) (int32, int64, int32, error)
}

type UserType interface {
	CreateUserType(userType domain.UserType) error
	GetUserType(name string) (domain.UserType, error)
}

type Fetter interface {
	CreateFetter(fetter domain.Fetter) error
	GetFetter(id int32) (domain.Fetter, error)
	GetFetterByName(name string) (domain.Fetter, error)
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

type Slave interface {
	CreateOrUpdateSlave(userId int32, masterId int32) error
	GetMaster(userId int32) (int32, error)
	GetSlaves(userId int32) ([]domain.SlavesListInfo, error)
}

type Repository struct {
	Authorization Authorization
	UserType      UserType
	Fetter        Fetter
	SlaveLevel    SlaveLevel
	SlaveStats    SlaveStats
	DefenderLevel DefenderLevel
	DefenderStats DefenderStats
	Slave         Slave
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
		Slave:         NewSlavePostgres(db),
	}
}
