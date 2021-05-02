package service

import (
	"errors"
	"math"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jackc/pgx/v4"
)

type AuthService struct {
	repAuth  repository.Authorization
	repSlave repository.Slave
}

func NewAuthService(repAuth repository.Authorization, repSlave repository.Slave) *AuthService {
	return &AuthService{
		repAuth:  repAuth,
		repSlave: repSlave,
	}
}

func (serv *AuthService) GetUser(id int32) (domain.User, error) {
	user, err := serv.repAuth.GetUser(id)

	return user, err
}

func (serv *AuthService) CreateUser(user domain.User) error {
	err := serv.repAuth.CreateUser(user)

	return err
}

func (serv *AuthService) GetUserVkInfo(token string) (domain.UserVkInfo, error) {
	vk := api.NewVK(token)
	res, err := vk.UsersGet(api.Params{
		"fields": "screen_name, photo_100",
	})

	if err != nil {
		return domain.UserVkInfo{}, err
	}

	us := res[0]

	return domain.UserVkInfo{
		Id:        int32(us.ID),
		Firstname: us.FirstName,
		Lastname:  us.LastName,
		IsClosed:  bool(us.IsClosed),
		Username:  us.ScreenName,
		Photo:     us.Photo100,
	}, nil
}

func (serv *AuthService) GetFriendsList(token string, friendId int32) ([]domain.FriendInfo, error) {
	vk := api.NewVK(token)

	res, err := vk.AppsGetFriendsListExtended(api.Params{
		"fields": "screen_name, photo_100",
	})
	if err != nil {
		return []domain.FriendInfo{}, err
	}

	friends := make([]domain.FriendInfo, res.Count)

	for i, fr := range res.Items {
		frInfLoc, err := serv.GetFriendInfoLocal(friendId)
		if err != nil {
			return friends, err
		}

		if frInfLoc.MasterId != 0 {
			res, err := vk.UsersGet(api.Params{
				"fields":  "screen_name, photo_100",
				"user_id": frInfLoc.MasterId,
			})

			if err != nil {
				return friends, err
			}

			us := res[0]

			frInfLoc.MasterFirstname = us.FirstName
			frInfLoc.MasterLastname = us.LastName
		}

		friends[i] = domain.FriendInfo{
			Id:          int32(fr.ID),
			Firstname:   fr.FirstName,
			Lastname:    fr.LastName,
			Photo:       fr.Photo100,
			FrInfoLocal: &frInfLoc,
		}
	}

	return friends, nil
}

func (serv *AuthService) GetFriendInfoLocal(friendId int32) (domain.FriendInfoLocal, error) {
	frInfLoc, err := serv.repAuth.GetFriendInfoLocal(friendId)

	if err.Error() == "no rows in result set" {
		return domain.FriendInfoLocal{
			MasterId:        0,
			MasterFirstname: "",
			MasterLastname:  "",
			HasFetter:       false,
			FetterType:      "common",
			PurchasePriceSm: 20,
			PurchasePriceGm: 0,
			SlaveLevel:      0,
			DefenderLevel:   0,
		}, nil
	}

	return frInfLoc, err
}

func (serv *AuthService) BuySlave(userId int32, slaveId int32) error {
	if userId == slaveId {
		return errors.New("Can't buy yourself")
	}

	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return err
	}

	slave, err := serv.repAuth.GetUser(slaveId)
	if err != nil {
		return err
	}

	if slave.HasFetter {
		if int32(time.Now().Add(3*time.Hour).UTC().Sub(slave.FetterTime).Minutes()) > user.FetterType.Duration {
			slave.HasFetter = false
			serv.repAuth.SetHasFetter(slave.Id, false)
		} else {
			return errors.New("Slave has fetter, you can't buy him")
		}
	}

	if user.Balance < slave.PurchasePriceSm || user.Gold < slave.PurchasePriceGm {
		return errors.New("Not enough money to buy a slave")
	}

	masterId, err := serv.repSlave.GetMaster(slaveId)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if masterId != 0 {
		if masterId == userId {
			return errors.New("Can't buy your slave")
		} else {
			slavesCount, balance, gold, err := serv.repAuth.GetUserBalance(masterId)
			if err != nil {
				return err
			}

			if err := serv.repAuth.SlaveCountBalanceUpdate(
				masterId,
				slavesCount-1,
				balance+slave.PurchasePriceSm,
				gold+slave.PurchasePriceGm); err != nil {
				return err
			}
		}
	}

	if err := serv.repAuth.SlaveCountBalanceUpdate(
		userId,
		user.SlavesCount+1,
		user.Balance-slave.PurchasePriceSm,
		user.Gold-slave.PurchasePriceGm); err != nil {
		return err
	}

	if err := serv.repSlave.CreateOrUpdateSlave(slaveId, userId); err != nil {
		return err
	}

	if err := serv.repAuth.SlaveBuyUpdateInfo(domain.SlaveBuyUpdateInfo{
		SlaveId:         slaveId,
		JobName:         "",
		UserType:        "slave",
		PurchasePriceSm: int64(math.Round(float64(slave.PurchasePriceSm) * 1.2)),
		SalePriceSm:     int64(math.Round(float64(slave.PurchasePriceSm) * 0.8)),
	}); err != nil {
		return err
	}

	return nil
}
