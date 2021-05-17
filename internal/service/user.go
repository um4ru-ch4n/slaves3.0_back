package service

import (
	"math"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const RATING_LIMIT = 100

type AuthService struct {
	repAuth       repository.User
	repUserMaster repository.UserMaster
}

func NewAuthService(repAuth repository.User, repUserMaster repository.UserMaster) *AuthService {
	return &AuthService{
		repAuth:       repAuth,
		repUserMaster: repUserMaster,
	}
}

func (serv *AuthService) GetUser(id int32) (domain.User, error) {
	user, err := serv.repAuth.GetUser(id)

	return user, errors.Wrap(err, "GetUser serv.repAuth.GetUser AuthService")
}

func (serv *AuthService) CreateUser(userId int32, userType, fio, photo string) (domain.UserFull, error) {

	user, err := serv.repAuth.CreateUser(userId, userType, fio, photo)
	if err != nil {
		return domain.UserFull{}, errors.Wrap(err, "CreateUser serv.repAuth.CreateUser AuthService")
	}

	userFull, err := serv.setAddFields(user)

	return userFull, errors.Wrap(err, "CreateUser serv.setAddFields AuthService")
}

func (serv *AuthService) GetUserVkInfo(token string) (domain.UserVkInfo, error) {
	vk := api.NewVK(token)
	res, err := vk.UsersGet(api.Params{
		"fields": "photo_100",
	})

	if err != nil {
		return domain.UserVkInfo{}, errors.Wrap(err, "GetUserVkInfo vk.UsersGet AuthService")
	}

	return domain.UserVkInfo{
		Id:    int32(res[0].ID),
		Fio:   res[0].LastName + " " + res[0].FirstName,
		Photo: res[0].Photo100,
	}, nil
}

func (serv *AuthService) GetUsersVkInfo(token string, usersIds []int32) ([]domain.UserVkInfo, error) {
	vk := api.NewVK(token)
	res, err := vk.UsersGet(api.Params{
		"fields":   "photo_100",
		"user_ids": usersIds,
	})

	if err != nil {
		return []domain.UserVkInfo{}, errors.Wrap(err, "GetUsersVkInfo vk.UsersGet AuthService")
	}

	users := make([]domain.UserVkInfo, len(usersIds))

	for i, _ := range res {
		users[i] = domain.UserVkInfo{
			Id:    int32(res[i].ID),
			Fio:   res[i].LastName + " " + res[i].FirstName,
			Photo: res[i].Photo100,
		}
	}

	return users, nil
}

func (serv *AuthService) GetUserIncome(userId int32) (int64, error) {
	slaves, err := serv.repUserMaster.GetSlaves(userId)
	if err != nil {
		return 0, errors.Wrap(err, "GetUserIncome serv.repUserMaster.GetSlaves AuthService")
	}

	var income int64

	for i, _ := range slaves {
		income += int64(GetSlaveProfit(slaves[i].SlaveLevel))
	}

	return income, nil
}

func (serv *AuthService) GetUserFull(id int32) (domain.UserFull, error) {
	user, err := serv.repAuth.GetUser(id)
	if err != nil {
		return domain.UserFull{}, errors.Wrap(err, "GetUserFull serv.repAuth.GetUser AuthService")
	}

	userFull, err := serv.setAddFields(user)

	return userFull, err
}

func (serv *AuthService) setAddFields(user domain.User) (domain.UserFull, error) {
	slaves, err := serv.repUserMaster.GetSlaves(user.Id)
	if err != nil {
		return domain.UserFull{}, errors.Wrap(err, "setAddFields  serv.repUserMaster.GetSlaves AuthService")
	}

	slavesCount := int32(len(slaves))

	var income int64

	for i, _ := range slaves {
		income += int64(GetSlaveProfit(slaves[i].SlaveLevel))
	}

	profit := GetSlaveProfit(user.SlaveLevel)
	damage := GetDefenderDamage(user.DefenderLevel)

	masterId, err := serv.repUserMaster.GetMaster(user.Id)
	if err != nil && err != pgx.ErrNoRows {
		return domain.UserFull{}, errors.Wrap(err, "setAddFields serv.repUserMaster.GetMaster AuthService")
	}

	userFull := domain.UserFull{
		Id:              user.Id,
		MasterId:        masterId,
		Fio:             user.Fio,
		Photo:           user.Photo,
		Balance:         user.Balance,
		Gold:            user.Gold,
		SlavesCount:     slavesCount,
		Income:          income,
		Profit:          profit,
		MoneyToUpdate:   GetSlaveMoneyToUpdate(user.SlaveLevel),
		Hp:              GetDefenderHp(user.DefenderLevel),
		Damage:          damage,
		DamageToUpdate:  GetDefenderDamageToUpdate(user.DefenderLevel),
		LastUpdate:      user.LastUpdate,
		JobName:         user.JobName,
		UserType:        user.UserType,
		SlaveLevel:      user.SlaveLevel,
		MoneyQuantity:   user.MoneyQuantity,
		DefenderLevel:   user.DefenderLevel,
		DamageQuantity:  user.DamageQuantity,
		PurchasePriceSm: GetUserPurchasePriceSm(user.SlaveLevel),
		SalePriceSm:     GetUserSalePriceSm(user.SlaveLevel),
		PurchasePriceGm: GetUserPurchasePriceGm(user.DefenderLevel),
		SalePriceGm:     GetUserSalePriceGm(user.DefenderLevel),
		HasFetter:       GetHasFetter(user.FetterTime, user.FetterDuration),
		FetterTime:      user.FetterTime,
		FetterType:      user.FetterType,
		FetterPrice:     user.FetterPrice,
		FetterDuration:  user.FetterDuration,
	}

	return userFull, nil
}

func (serv *AuthService) GetFriendsList(token string) ([]domain.FriendInfo, error) {
	vk := api.NewVK(token)

	ids, err := vk.FriendsGet(api.Params{})
	if err != nil {
		return []domain.FriendInfo{}, errors.Wrap(err, "GetFriendsList vk.FriendsGet AuthService")
	}

	res, err := vk.UsersGet(api.Params{
		"fields":   "photo_100",
		"user_ids": ids.Items,
	})
	if err != nil {
		return []domain.FriendInfo{}, errors.Wrap(err, "GetFriendsList vk.UsersGet AuthService")
	}

	friends, err := serv.repAuth.GetFriendsInfo(ids.Items)
	if err != nil {
		return []domain.FriendInfo{}, errors.Wrap(err, "GetFriendsList serv.repAuth.GetFriendsInfo AuthService")
	}

	mastersIds := make([]int32, 0, len(friends))

	for i, _ := range friends {
		if friends[i].MasterId != 0 {
			mastersIds = append(mastersIds, friends[i].MasterId)
		}
	}

	masters := make(map[int32]string, len(mastersIds))

	if len(mastersIds) > 0 {
		mst, err := vk.UsersGet(api.Params{
			"user_ids": mastersIds,
		})

		if err != nil {
			return []domain.FriendInfo{}, errors.Wrap(err, "GetFriendsList vk.UsersGet AuthService")
		}

		for i, _ := range mst {
			masters[int32(mst[i].ID)] = mst[i].LastName + " " + mst[i].FirstName
		}
	}

	friendsInfo := make([]domain.FriendInfo, 0, 100)

	for i, _ := range res {
		friendsInfo = append(friendsInfo, domain.FriendInfo{
			Id:    int32(res[i].ID),
			Fio:   res[i].LastName + " " + res[i].FirstName,
			Photo: res[i].Photo100,
		})

		if val, ok := friends[int32(res[i].ID)]; ok {
			friendsInfo[i].HasFetter = GetHasFetter(val.FetterTime, val.FetterDuration)
			friendsInfo[i].MasterId = val.MasterId
			friendsInfo[i].FetterTime = val.FetterTime
			friendsInfo[i].FetterType = val.FetterType
			friendsInfo[i].FetterDuration = val.FetterDuration
			friendsInfo[i].PurchasePriceSm = GetUserPurchasePriceSm(val.SlaveLevel)
			friendsInfo[i].PurchasePriceGm = GetUserPurchasePriceGm(val.DefenderLevel)
			friendsInfo[i].SlaveLevel = val.SlaveLevel
			friendsInfo[i].DefenderLevel = val.DefenderLevel

			if val.MasterId != 0 {
				friendsInfo[i].MasterFIO = masters[val.MasterId]
			}
		} else {
			friendsInfo[i].HasFetter = false
			friendsInfo[i].MasterId = 0
			friendsInfo[i].FetterType = "common"
			friendsInfo[i].FetterDuration = 120
			friendsInfo[i].PurchasePriceSm = GetUserPurchasePriceSm(1)
			friendsInfo[i].PurchasePriceGm = GetUserPurchasePriceGm(1)
			friendsInfo[i].SlaveLevel = 1
			friendsInfo[i].DefenderLevel = 1
		}
	}

	return friendsInfo, nil
}

func (serv *AuthService) BuySlave(userId int32, slaveId int32) error {
	if userId == slaveId {
		return errors.Wrap(errors.New("Can't buy yourself"), "BuySlave userId == slaveId AuthService")
	}

	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return errors.Wrap(err, "BuySlave get user info serv.repAuth.GetUser AuthService")
	}

	slave, err := serv.repAuth.GetUser(slaveId)
	if err != nil {
		return errors.Wrap(err, "BuySlave get slave info serv.repAuth.GetUser AuthService")
	}

	slaveHasFetter := GetHasFetter(slave.FetterTime, slave.FetterDuration)

	if slaveHasFetter {
		return errors.Wrap(errors.New("Slave has fetter, you can't buy him"), "BuySlave slave has fetter AuthService")
	}

	slavePurchasePriceSm := GetUserPurchasePriceSm(slave.SlaveLevel)
	slavePurchasePriceGm := GetUserPurchasePriceGm(slave.DefenderLevel)

	if user.Balance < slavePurchasePriceSm || user.Gold < slavePurchasePriceGm {
		return errors.Wrap(errors.New("Not enough money to buy a slave"), "BuySlave user has enough money AuthService")
	}

	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "BuySlave serv.repUserMaster.GetMaster AuthService")
	}

	if masterId != 0 {
		if masterId == userId {
			return errors.Wrap(errors.New("Can't buy your slave"), "BuySlave masterId == userID AuthService")
		} else {
			balance, gold, err := serv.repAuth.GetUserBalance(masterId)
			if err != nil {
				return errors.Wrap(err, "BuySlave serv.repAuth.GetUserBalance AuthService")
			}

			if err := serv.repAuth.UserBalanceUpdate(
				masterId,
				balance+slavePurchasePriceSm,
				gold+slavePurchasePriceGm); err != nil {
				return errors.Wrap(err, "BuySlave serv.repAuth.UserBalanceUpdate master AuthService")
			}
		}
	}

	if err := serv.repAuth.UserBalanceUpdate(
		userId,
		user.Balance-slavePurchasePriceSm,
		user.Gold-slavePurchasePriceGm); err != nil {
		return errors.Wrap(err, "BuySlave serv.repAuth.UserBalanceUpdate user AuthService")
	}

	if err := serv.repUserMaster.CreateOrUpdateSlave(slaveId, userId); err != nil {
		return errors.Wrap(err, "BuySlave serv.repUserMaster.CreateOrUpdateSlave AuthService")
	}

	if err := serv.repAuth.SlaveBuyUpdateInfo(domain.SlaveBuyUpdateInfo{
		SlaveId:  slaveId,
		JobName:  "",
		UserType: "slave",
	}); err != nil {
		return errors.Wrap(err, "BuySlave serv.repAuth.SlaveBuyUpdateInfo AuthService")
	}

	return nil
}

func (serv *AuthService) SaleSlave(userId int32, slaveId int32) error {
	if userId == slaveId {
		return errors.Wrap(errors.New("Can't sale yourself"), "SaleSlave userId == slaveId AuthService")
	}

	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return errors.Wrap(err, "SaleSlave serv.repAuth.GetUser user AuthService")
	}

	slave, err := serv.repAuth.GetUser(slaveId)
	if err != nil {
		return errors.Wrap(err, "SaleSlave serv.repAuth.GetUser slave AuthService")
	}

	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "SaleSlave serv.repUserMaster.GetMaster AuthService")
	}

	if masterId == 0 {
		return errors.Wrap(errors.New("Can't sale free slave"), "SaleSlave masterId == 0 slave AuthService")
	}

	if masterId != userId {
		return errors.Wrap(err, "SaleSlave masterId != userId AuthService")
	}

	if err := serv.repAuth.UserBalanceUpdate(
		userId,
		user.Balance+GetUserSalePriceSm(slave.SlaveLevel),
		user.Gold+GetUserSalePriceGm(slave.DefenderLevel)); err != nil {
		return errors.Wrap(err, "SaleSlave serv.repAuth.UserBalanceUpdate AuthService")
	}

	if err := serv.repUserMaster.SaleSlave(slaveId); err != nil {
		return errors.Wrap(err, "SaleSlave serv.repUserMaster.SaleSlave AuthService")
	}

	if err := serv.repAuth.SlaveBuyUpdateInfo(domain.SlaveBuyUpdateInfo{
		SlaveId:  slaveId,
		JobName:  "",
		UserType: "simp",
	}); err != nil {
		return errors.Wrap(err, "SaleSlave serv.repAuth.SlaveBuyUpdateInfo AuthService")
	}

	return nil
}

func (serv *AuthService) GetRatingBySlavesCount() ([]domain.RatingSlavesCount, error) {
	ratingSlavesCount, err := serv.repAuth.GetRatingBySlavesCount(RATING_LIMIT)
	if err != nil {
		return ratingSlavesCount, errors.Wrap(err, "GetRatingBySlavesCount serv.repAuth.GetRatingBySlavesCount AuthService")
	}

	for i, _ := range ratingSlavesCount {
		ratingSlavesCount[i].HasFetter = GetHasFetter(
			ratingSlavesCount[i].FetterTime,
			ratingSlavesCount[i].FetterDuration)
	}

	return ratingSlavesCount, nil
}

func (serv *AuthService) GetSlavesList(userId int32) ([]domain.SlavesListInfo, error) {
	slavesList, err := serv.repUserMaster.GetSlaves(userId)
	if err != nil {
		return slavesList, errors.Wrap(err, "GetSlavesList serv.repUserMaster.GetSlaves AuthService")
	}

	for i, _ := range slavesList {
		slavesList[i].HasFetter = GetHasFetter(slavesList[i].FetterTime, slavesList[i].FetterDuration)
		slavesList[i].Profit = int64(GetSlaveProfit(slavesList[i].SlaveLevel))
	}

	return slavesList, nil
}

func (serv *AuthService) SetJobName(userId, slaveId int32, jobName string) error {
	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil {
		return errors.Wrap(err, "SetJobName serv.repUserMaster.GetMaster AuthService")
	}

	if masterId != userId {
		return errors.Wrap(errors.New("You can't set job name for other's slave"), "SetJobName masterId != userId AuthService")
	}

	err = serv.repAuth.SetJobName(slaveId, jobName)

	return errors.Wrap(err, "SetJobName serv.repAuth.SetJobName AuthService")
}

func (serv *AuthService) GetLastUpdate(userId int32) (time.Time, error) {
	time, err := serv.repAuth.GetLastUpdate(userId)
	if err != nil {
		return time, errors.Wrap(err, "GetLastUpdate serv.repAuth.GetLastUpdate AuthService")
	}

	return time, nil
}

func (serv *AuthService) UpdateUserInfo(userId int32) error {
	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return errors.Wrap(err, "UpdateUserInfo serv.repAuth.GetUser AuthService")
	}

	slaves, err := serv.repUserMaster.GetSlavesForUpdate(userId)
	if err != nil {
		return errors.Wrap(err, "UpdateUserInfo serv.repUserMaster.GetSlavesForUpdate AuthService")
	}

	var balancePerTime, slaveMoneyToUpdate, incBalance int64
	var slaveProfit int32
	var tmpTime float64

	timeSinceLUpd := time.Since(user.LastUpdate).Minutes()
	timeSinceCopy := timeSinceLUpd

	for i, _ := range slaves {
		slaveProfit = GetSlaveProfit(slaves[i].SlaveLevel)
		balancePerTime = int64(math.Round(float64(slaveProfit) * timeSinceLUpd))
		slaveMoneyToUpdate = GetSlaveMoneyToUpdate(slaves[i].SlaveLevel)

		for balancePerTime+slaves[i].MoneyQuantity >= slaveMoneyToUpdate {
			tmpTime = float64((slaveMoneyToUpdate - slaves[i].MoneyQuantity) / int64(slaveProfit))

			timeSinceCopy -= tmpTime
			incBalance += (slaveMoneyToUpdate - slaves[i].MoneyQuantity)

			slaves[i].SlaveLevel++
			slaveProfit = GetSlaveProfit(slaves[i].SlaveLevel)
			balancePerTime = int64(math.Round(float64(slaveProfit) * timeSinceCopy))
			slaveMoneyToUpdate = GetSlaveMoneyToUpdate(slaves[i].SlaveLevel)

			slaves[i].MoneyQuantity = 0

		}
		incBalance += int64(math.Round(timeSinceCopy * float64(slaveProfit)))
		slaves[i].MoneyQuantity += int64(math.Round(timeSinceCopy * float64(slaveProfit)))
		timeSinceCopy = timeSinceLUpd

		if err := serv.repAuth.UpdateSlaveHour(slaves[i]); err != nil {
			return errors.Wrap(err, "UpdateUserInfo serv.repAuth.UpdateSlaveHour AuthService")
		}
	}

	user.Balance += incBalance
	user.LastUpdate = time.Now()

	serv.repAuth.UpdateUserBalanceHour(userId, user.Balance)

	return nil
}
