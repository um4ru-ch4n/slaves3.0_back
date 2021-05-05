package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
)

type User interface {
	CreateUser(userId int32, userType string) (domain.UserFull, error)
	GetUserFull(id int32) (domain.UserFull, error)
	GetUserVkInfo(token string) (domain.UserVkInfo, error)
	GetFriendsList(token string) ([]domain.FriendInfo, error)
	BuySlave(userId int32, slaveId int32) error
	SaleSlave(userId int32, slaveId int32) error
}

type Service struct {
	User User
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		User: NewAuthService(rep.User, rep.UserMaster),
	}
}
