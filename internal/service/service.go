package service

import (
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type User interface {
	CreateUser(userId int32, userType, fio, photo string) (domain.UserFull, error)
	GetUserVkInfo(token string) (domain.UserVkInfo, error)
	GetUsersVkInfo(token string, usersIds []int32) ([]domain.UserVkInfo, error)
	GetUserFull(id int32) (domain.UserFull, error)
	GetFriendsList(token string) ([]domain.FriendInfo, error)
	BuySlave(userId int32, slaveId int32) error
	SaleSlave(userId int32, slaveId int32) error
	GetRatingBySlavesCount() ([]domain.RatingSlavesCount, error)
	GetSlavesList(userId int32) ([]domain.SlavesListInfo, error)
	SetJobName(userId, slaveId int32, jobName string) error
	GetUser(id int32) (domain.User, error)
	GetUserIncome(userId int32) (int64, error)
	GetLastUpdate(userId int32) (time.Time, error)
	UpdateUserInfo(userId int32) error
}

type Service struct {
	User User
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		User: NewAuthService(rep.User, rep.UserMaster),
	}
}
