package service

import (
	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/SevereCloud/vksdk/v2/api"
)

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{
		rep: rep,
	}
}

func (serv *AuthService) GetUser(id int32) (domain.User, error) {
	user, err := serv.rep.GetUser(id)

	return user, err
}

func (serv *AuthService) CreateUser(user domain.User) error {
	err := serv.rep.CreateUser(user)

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
	frInfLoc, err := serv.rep.GetFriendInfoLocal(friendId)

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
