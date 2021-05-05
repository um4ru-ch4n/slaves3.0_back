package service

import (
	"math"
	"time"
)

func GetSlaveProfit(slaveLevel int32) int32 {
	return 1
}

func GetSlaveMoneyToUpdate(slaveLevel, slaveProfit int32) int64 {
	return 10
}

func GetDefenderHp(defenderLevel int32) int32 {
	return 5
}

func GetDefenderDamage(defenderLevel int32) int32 {
	return 1
}

func GetDefenderDamageToUpdate(defenderLevel, defenderDamage int32) int64 {
	return 10
}

func GetUserSalePriceSm(purchasePriceSm int64) int64 {
	return int64(math.Round(float64(purchasePriceSm) * 0.8))
}

func GetUserSalePriceGm(purchasePriceGm int32) int32 {
	return int32(math.Round(float64(purchasePriceGm) * 0.8))
}

func GetHasFetter(fetterTime time.Time, duration int32) bool {
	if fetterTime.Year() == 1971 {
		return false
	}

	if int32(time.Now().Sub(fetterTime).Minutes()) > duration {
		return false
	}
	return true
}

func IncSlavePurchasePriceSm(purchasePriceSm int64) int64 {
	return int64(math.Round(float64(purchasePriceSm) * 1.2))
}
