package repository

import (
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type User interface {
	CreateUser(userId int32, userType, fio, photo string) (domain.User, error)
	GetUser(id int32) (domain.User, error)
	GetUserType(userId int32) (string, error)
	GetFriendsInfo(ids []int) (map[int32]domain.FriendInfo, error)
	SlaveBuyUpdateInfo(newData domain.SlaveBuyUpdateInfo) error
	UserBalanceUpdate(userId int32, balance int64, gold int32) error
	SetFetterTime(userId int32, fetterTime time.Time) error
	GetFetterTime(userId int32) (time.Time, error)
	GetUserBalance(userId int32) (int64, int32, error)
	GetRatingBySlavesCount(limit int32) ([]domain.RatingSlavesCount, error)
	SetJobName(slaveId int32, jobName string) error
	GetLastUpdate(userId int32) (time.Time, error)
	UpdateUserBalanceHour(userId int32, balance int64) error
	UpdateSlaveHour(slaveInfo domain.SlaveInfoForUpdate) error
}

type UserMaster interface {
	CreateOrUpdateSlave(userId int32, masterId int32) error
	GetMaster(userId int32) (int32, error)
	GetSlaves(userId int32) ([]domain.SlavesListInfo, error)
	SaleSlave(slaveId int32) error
	GetSlavesForUpdate(userId int32) ([]domain.SlaveInfoForUpdate, error)
}

type Repository struct {
	User       User
	UserMaster UserMaster
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		User:       NewAuthPostgres(db),
		UserMaster: NewUserMasterPostgres(db),
	}
}
