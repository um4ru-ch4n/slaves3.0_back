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
		FirstName: us.FirstName,
		LastName:  us.LastName,
		IsClosed:  bool(us.IsClosed),
		Username:  us.ScreenName,
		Photo:     us.Photo100,
	}, nil
}

// func (serv *AuthService) createUser(user *domain.User) error {
// 	userType := domain.UserType{Name: "default"}
// 	err := rep.db.Model(&userType).Where("name = ?", userType.Name).Select()

// 	if err != nil {
// 		return user, err
// 	}

// 	slavesLevel := domain.SlaveLevel{Lvl: 0}

// 	err = rep.db.Model(&slavesLevel).Where("lvl = ?", slavesLevel.Id).Select()

// 	if err != nil {
// 		return user, err
// 	}

// 	slavesStats := domain.SlaveStats{
// 		Level:          slavesLevel,
// 		Money_quantity: 0,
// 	}

// 	defenderLevel := domain.DefenderLevel{Lvl: 0}

// 	err = rep.db.Model(&defenderLevel).Where("lvl = ?", defenderLevel.Id).Select()

// 	if err != nil {
// 		return user, err
// 	}

// 	defenderStats := domain.DefenderStats{
// 		Level:           defenderLevel,
// 		Damage_quantity: 0,
// 	}

// 	fetterType := domain.Fetter{Name: "common"}

// 	err = rep.db.Model(&fetterType).Where("name = ?", fetterType.Name).Select()

// 	if err != nil {
// 		return user, err
// 	}

// 	user = domain.User{
// 		Id:                    id,
// 		Slaves_count:          0,
// 		Balance:               100,
// 		Income:                0,
// 		Last_update:           time.Now(),
// 		Job_name:              "",
// 		User_type:             userType,
// 		Slave_stats:           slavesStats,
// 		Defender_stats:        defenderStats,
// 		Purchase_price_silver: 10,
// 		Sale_price_silver:     5,
// 		Purchase_price_gold:   0,
// 		Sale_price_gold:       0,
// 		Has_fetter:            false,
// 		Fetter_time:           time.Now(),
// 		Fetter_type:           fetterType,
// 		Slaves:                []domain.User{},
// 	}

// 	_, err = rep.db.Model(&user).Insert()

// 	if err != nil {
// 		return user, err
// 	}
// }
