package service

import (
	"math"
	"time"
)

func GetSlaveProfit(slaveLevel int32) int32 {
	slLvlFl := float64(slaveLevel)
	return int32(math.Round(math.Pow(math.Log(slLvlFl), 3)*math.Pow(slLvlFl, 2)/2 + 1))
}

func GetSlaveMoneyToUpdate(slaveLevel int32) int64 {
	slLvlFl := float64(slaveLevel)
	return int64(math.Round(math.Pow(math.Log(slLvlFl), 3)*math.Pow(slLvlFl, 4) + 10))
}

func GetDefenderHp(defenderLevel int32) int32 {
	defLvlFl := float64(defenderLevel)
	return int32(math.Round(math.Log(defLvlFl)*100*defLvlFl + 100))
}

func GetDefenderDamage(defenderLevel int32) int32 {
	defLvlFl := float64(defenderLevel)
	return int32(math.Round((math.Log(defLvlFl)*100*defLvlFl + 100) * 0.1))
}

func GetDefenderDamageToUpdate(defenderLevel int32) int64 {
	defLvlFl := float64(defenderLevel)
	return int64(10 * defenderLevel * int32(math.Round(math.Log(defLvlFl)*100*defLvlFl+100)))
}

func GetUserPurchasePriceSm(slaveLevel int32) int64 {
	slLvlFl := float64(slaveLevel)
	return int64(math.Round(math.Pow(slLvlFl, 2)*(math.Pow(math.Log(slLvlFl), 3)*math.Pow(slLvlFl, 2)/2+1) + 1))
}

func GetUserPurchasePriceGm(defenderLevel int32) int32 {
	if defenderLevel == 1 {
		return 0
	}
	defLvlFl := float64(defenderLevel)
	return int32(math.Round(defLvlFl * math.Pow(1.1, defLvlFl-2)))
}

func GetUserSalePriceSm(slaveLevel int32) int64 {
	if slaveLevel == 1 {
		return 1
	}
	slLvlFl := float64(slaveLevel - 1)
	return int64(math.Round((math.Pow(slLvlFl, 2)*(math.Pow(math.Log(slLvlFl), 3)*math.Pow(slLvlFl, 2)/2+1) + 1) * 0.8))
}

func GetUserSalePriceGm(defenderLevel int32) int32 {
	if defenderLevel == 1 {
		return 0
	}
	defLvlFl := float64(defenderLevel - 1)
	return int32(math.Round((defLvlFl * math.Pow(1.1, defLvlFl-2) * 0.8)))
}

func GetHasFetter(fetterTime time.Time, duration int32) bool {
	return int32(time.Since(fetterTime).Minutes()) < duration
}
