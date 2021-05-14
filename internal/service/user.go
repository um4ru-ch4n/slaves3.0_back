package service

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/00mrx00/slaves3.0_back/internal/repository"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jackc/pgx/v4"
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

	return user, err
}

func (serv *AuthService) CreateUser(userId int32, userType, fio, photo string) (domain.UserFull, error) {

	user, err := serv.repAuth.CreateUser(userId, userType, fio, photo)
	if err != nil {
		return domain.UserFull{}, err
	}

	userFull, err := serv.setAddFields(user)

	return userFull, err
}

func (serv *AuthService) GetUserVkInfo(token string) (domain.UserVkInfo, error) {
	vk := api.NewVK(token)
	res, err := vk.UsersGet(api.Params{
		"fields": "photo_100",
	})

	if err != nil {
		return domain.UserVkInfo{}, err
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
		return []domain.UserVkInfo{}, err
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
		fmt.Println(err)
		return 0, err
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
		return domain.UserFull{}, err
	}

	userFull, err := serv.setAddFields(user)

	return userFull, err
}

func (serv *AuthService) setAddFields(user domain.User) (domain.UserFull, error) {
	slaves, err := serv.repUserMaster.GetSlaves(user.Id)
	if err != nil {
		fmt.Println(err)
		return domain.UserFull{}, err
	}

	slavesCount := int32(len(slaves))

	var income int64

	for i, _ := range slaves {
		income += int64(GetSlaveProfit(slaves[i].SlaveLevel))
	}

	profit := GetSlaveProfit(user.SlaveLevel)
	damage := GetDefenderDamage(user.DefenderLevel)

	userFull := domain.UserFull{
		Id:              user.Id,
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
		return []domain.FriendInfo{}, err
	}

	res, err := vk.UsersGet(api.Params{
		"fields":   "photo_100",
		"user_ids": ids.Items,
	})
	if err != nil {
		return []domain.FriendInfo{}, err
	}

	friendsIds := make([]int32, len(res))

	for i, _ := range res {
		friendsIds[i] = int32(res[i].ID)
	}

	friends, err := serv.repAuth.GetFriendsInfo(friendsIds)
	if err != nil {
		return []domain.FriendInfo{}, err
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
			return []domain.FriendInfo{}, err
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

	slaveHasFetter := GetHasFetter(slave.FetterTime, slave.FetterDuration)

	if slaveHasFetter {
		return errors.New("Slave has fetter, you can't buy him")
	}

	slavePurchasePriceSm := GetUserPurchasePriceSm(slave.SlaveLevel)
	slavePurchasePriceGm := GetUserPurchasePriceGm(slave.DefenderLevel)

	if user.Balance < slavePurchasePriceSm || user.Gold < slavePurchasePriceGm {
		return errors.New("Not enough money to buy a slave")
	}

	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if masterId != 0 {
		if masterId == userId {
			return errors.New("Can't buy your slave")
		} else {
			balance, gold, err := serv.repAuth.GetUserBalance(masterId)
			if err != nil {
				return err
			}

			if err := serv.repAuth.UserBalanceUpdate(
				masterId,
				balance+slavePurchasePriceSm,
				gold+slavePurchasePriceGm); err != nil {
				return err
			}
		}
	}

	if err := serv.repAuth.UserBalanceUpdate(
		userId,
		user.Balance-slavePurchasePriceSm,
		user.Gold-slavePurchasePriceGm); err != nil {
		return err
	}

	if err := serv.repUserMaster.CreateOrUpdateSlave(slaveId, userId); err != nil {
		return err
	}

	if err := serv.repAuth.SlaveBuyUpdateInfo(domain.SlaveBuyUpdateInfo{
		SlaveId:  slaveId,
		JobName:  "",
		UserType: "slave",
	}); err != nil {
		return err
	}

	return nil
}

func (serv *AuthService) SaleSlave(userId int32, slaveId int32) error {
	if userId == slaveId {
		return errors.New("Can't sale yourself")
	}

	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return err
	}

	slave, err := serv.repAuth.GetUser(slaveId)
	if err != nil {
		return err
	}

	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if masterId == 0 {
		return errors.New("Can't sale free slave")
	}

	if masterId != userId {
		return errors.New("Can't sale other's slave")
	}

	if err := serv.repAuth.UserBalanceUpdate(
		userId,
		user.Balance+GetUserPurchasePriceSm(slave.SlaveLevel),
		user.Gold+GetUserPurchasePriceGm(slave.DefenderLevel)); err != nil {
		return err
	}

	if err := serv.repUserMaster.SaleSlave(slaveId); err != nil {
		return err
	}

	if err := serv.repAuth.SlaveBuyUpdateInfo(domain.SlaveBuyUpdateInfo{
		SlaveId:  slaveId,
		JobName:  "",
		UserType: "simp",
	}); err != nil {
		return err
	}

	return nil
}

func (serv *AuthService) GetRatingBySlavesCount() ([]domain.RatingSlavesCount, error) {
	ratingSlavesCount, err := serv.repAuth.GetRatingBySlavesCount(RATING_LIMIT)
	if err != nil {
		return ratingSlavesCount, err
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

	for i, _ := range slavesList {
		slavesList[i].HasFetter = GetHasFetter(slavesList[i].FetterTime, slavesList[i].FetterDuration)
		slavesList[i].Profit = int64(GetSlaveProfit(slavesList[i].SlaveLevel))
	}

	return slavesList, err
}

func (serv *AuthService) SetJobName(userId, slaveId int32, jobName string) error {
	masterId, err := serv.repUserMaster.GetMaster(slaveId)
	if err != nil {
		return err
	}

	if masterId != userId {
		return errors.New("You can't set job name for other's slave")
	}

	err = serv.repAuth.SetJobName(slaveId, jobName)

	return err
}

func (serv *AuthService) GetLastUpdate(userId int32) (time.Time, error) {
	return serv.repAuth.GetLastUpdate(userId)
}

func (serv *AuthService) UpdateUserInfo(userId int32) error {
	user, err := serv.repAuth.GetUser(userId)
	if err != nil {
		return err
	}

	slaves, err := serv.repUserMaster.GetSlavesForUpdate(userId)
	if err != nil {
		return err
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
			return err
		}
	}

	user.Balance += incBalance
	user.LastUpdate = time.Now()

	serv.repAuth.UpdateUserBalanceHour(userId, user.Balance)

	return nil
}
