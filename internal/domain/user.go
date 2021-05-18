package domain

import (
	"time"
)

type UserType struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Fetter struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Duration int32  `json:"duration"`
}

type User struct {
	Id             int32     `json:"id"`
	Fio            string    `json:"fio"`
	Photo          string    `json:"photo"`
	Balance        int64     `json:"balance"`
	Gold           int32     `json:"gold"`
	LastUpdate     time.Time `json:"last_update"`
	JobName        string    `json:"job_name"`
	UserType       string    `json:"user_type"`
	SlaveLevel     int32     `json:"slave_level"`
	MoneyQuantity  int64     `json:"money_quantity"`
	DefenderLevel  int32     `json:"defender_level"`
	DamageQuantity int64     `json:"damage_quantity"`
	FetterTime     time.Time `json:"fetter_time"`
	FetterType     string    `json:"fetter_type"`
	FetterPrice    int64     `json:"fetter_price"`
	FetterDuration int32     `json:"fetter_duration"`
	MasterId       int32     `json:"master_id"`
	MasterFio      string    `json:"master_fio"`
}

type UserFull struct {
	Id              int32     `json:"id"`
	Fio             string    `json:"fio"`
	Photo           string    `json:"photo"`
	Balance         int64     `json:"balance"`
	Gold            int32     `json:"gold"`
	SlavesCount     int32     `json:"slaves_count"`     //
	Income          int64     `json:"income"`           //
	Profit          int32     `json:"profit"`           //
	MoneyToUpdate   int64     `json:"money_to_update"`  //
	Hp              int32     `json:"hp"`               //
	Damage          int32     `json:"damage"`           //
	DamageToUpdate  int64     `json:"damage_to_update"` //
	LastUpdate      time.Time `json:"last_update"`
	JobName         string    `json:"job_name"`
	UserType        string    `json:"user_type"`
	SlaveLevel      int32     `json:"slave_level"`
	MoneyQuantity   int64     `json:"money_quantity"`
	DefenderLevel   int32     `json:"defender_level"`
	DamageQuantity  int64     `json:"damage_quantity"`
	PurchasePriceSm int64     `json:"purchase_price_sm"` //
	SalePriceSm     int64     `json:"sale_price_sm"`     //
	PurchasePriceGm int32     `json:"purchase_price_gm"` //
	SalePriceGm     int32     `json:"sale_price_gm"`     //
	HasFetter       bool      `json:"has_fetter"`        //
	FetterTime      time.Time `json:"fetter_time"`
	FetterType      string    `json:"fetter_type"`
	FetterPrice     int64     `json:"fetter_price"`
	FetterDuration  int32     `json:"fetter_duration"`
	MasterId        int32     `json:"master_id"`
	MasterFio       string    `json:"master_fio"`
}

type UserMaster struct {
	UserId   *User
	MasterId *User
}

type UserVkInfo struct {
	Id    int32  `json:"id"`
	Fio   string `json:"fio"`
	Photo string `json:"photo"`
}

type FriendInfo struct {
	Id              int32     `json:"id"`
	Fio             string    `json:"fio"`
	Photo           string    `json:"photo"`
	MasterId        int32     `json:"master_id"`
	MasterFIO       string    `json:"master_fio"`
	HasFetter       bool      `json:"has_fetter"`
	FetterTime      time.Time `json:"fetter_time"`
	FetterType      string    `json:"fetter_type"`
	FetterDuration  int32     `json:"fetter_duration"`
	PurchasePriceSm int64     `json:"purchase_price_sm"`
	PurchasePriceGm int32     `json:"purchase_price_gm"`
	SlaveLevel      int32     `json:"slave_level"`
	DefenderLevel   int32     `json:"defender_level"`
}

type SlavesListInfo struct {
	Id             int32     `json:"id"`
	Fio            string    `json:"fio"`
	Photo          string    `json:"photo"`
	Profit         int64     `json:"profit"`
	JobName        string    `json:"job_name"`
	HasFetter      bool      `json:"has_fetter"`
	FetterTime     time.Time `json:"fetter_time"`
	FetterType     string    `json:"fetter_type"`
	FetterDuration int32     `json:"fetter_duration"`
	SlaveLevel     int32     `json:"slave_level"`
	DefenderLevel  int32     `json:"defender_level"`
}

type SlaveBuyUpdateInfo struct {
	SlaveId  int32
	JobName  string
	UserType string
}

type SlaveId struct {
	SlaveId int32 `json:"slave_id" binding:"required"`
}

type RatingSlavesCount struct {
	Id             int32     `json:"id"`
	Fio            string    `json:"fio"`
	SlavesCount    int32     `json:"slaves_count"`
	Photo          string    `json:"photo"`
	HasFetter      bool      `json:"has_fetter"`
	FetterTime     time.Time `json:"fetter_time"`
	FetterType     string    `json:"fetter_type"`
	FetterDuration int32     `json:"fetter_duration"`
}

type SlaveInfoForUpdate struct {
	Id            int32  `json:"id"`
	UserType      string `json:"user_type"`
	SlaveLevel    int32  `json:"slave_level"`
	MoneyQuantity int64  `json:"money_quantity"`
}

type FetterBuySlaveInfo struct {
	FetterTime     time.Time `json:"fetter_time"`
	FetterDuration int32     `json:"fetter_duration"`
	FetterPrice    int32     `json:"fetter_price"`
	SlaveLevel     int32     `json:"slave_level"`
	DefenderLevel  int32     `json:"defender_level"`
}
