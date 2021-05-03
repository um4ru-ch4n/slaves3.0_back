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

// SlavesCount считаем запросом COUNT из таблицы user_master
// Income считаем через INNER JOIN user_master ON users по полю profit
// Profit and MoneyToUpdate calculates from SlaveLevel
// Hp, Damage, DamageToUpdate calculates from DefenderLevel
// SalePriceSm calculates from PurchasePriceSm
// SalePriceGm calculates from PurchasePriceGm
// HasFetter calculates from FetterTime. If FetterTime year = 1977, hasFetter = false otherwise true
type User struct {
	Id              int32       `json:"id"`
	Balance         int64       `json:"balance"`
	Gold            int32       `json:"gold"`
	LastUpdate      time.Time   `json:"last_update"`
	JobName         string      `json:"job_name"`
	UserType        string      `json:"user_type"`
	SlaveLevel      int32       `json:"slave_level"`
	MoneyQuantity   int64       `json:"money_quantity"`
	DefenderLevel   int32       `json:"defender_level"`
	DamageQuantity  int64       `json:"damage_quantity"`
	PurchasePriceSm int64       `json:"purchase_price_sm"`
	PurchasePriceGm int32       `json:"purchase_price_gm"`
	FetterTime      time.Time   `json:"fetter_time"`
	FetterType      *Fetter     `json:"fetter_type"`
	VkInfo          *UserVkInfo `json:"vk_info"`
}

type UserFull struct {
	Id              int32       `json:"id"`
	Balance         int64       `json:"balance"`
	Gold            int32       `json:"gold"`
	SlavesCount     int32       `json:"slaves_count"`     //
	Income          int64       `json:"income"`           //
	Profit          int32       `json:"profit"`           //
	MoneyToUpdate   int64       `json:"money_to_update"`  //
	Hp              int32       `json:"hp"`               //
	Damage          int32       `json:"damage"`           //
	DamageToUpdate  int64       `json:"damage_to_update"` //
	LastUpdate      time.Time   `json:"last_update"`
	JobName         string      `json:"job_name"`
	UserType        string      `json:"user_type"`
	SlaveLevel      int32       `json:"slave_level"`
	MoneyQuantity   int64       `json:"money_quantity"`
	DefenderLevel   int32       `json:"defender_level"`
	DamageQuantity  int64       `json:"damage_quantity"`
	PurchasePriceSm int64       `json:"purchase_price_sm"`
	SalePriceSm     int64       `json:"sale_price_sm"` //
	PurchasePriceGm int32       `json:"purchase_price_gm"`
	SalePriceGm     int32       `json:"sale_price_gm"` //
	HasFetter       bool        `json:"has_fetter"`    //
	FetterTime      time.Time   `json:"fetter_time"`
	FetterType      *Fetter     `json:"fetter_type"`
	VkInfo          *UserVkInfo `json:"vk_info"`
}

type UserMaster struct {
	UserId   *User
	MasterId *User
}

type UserVkInfo struct {
	Id        int32  `json:"id"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	IsClosed  bool   `json:"is_closed"`
	Username  string `json:"username"`
	Photo     string `json:"photo"`
}

type FriendInfoLocal struct {
	MasterId        int32     `json:"master_id"`
	MasterFirstname string    `json:"master_Firstname"`
	MasterLastname  string    `json:"master_Lastname"`
	FetterTime      time.Time `json:"fetter_time"`
	FetterType      string    `json:"fetter_type"`
	PurchasePriceSm int64     `json:"purchase_price_sm"`
	PurchasePriceGm int32     `json:"purchase_price_gm"`
	SlaveLevel      int32     `json:"slave_level"`
	DefenderLevel   int32     `json:"defender_level"`
}

type FriendInfo struct {
	Id          int32            `json:"id"`
	Firstname   string           `json:"Firstname"`
	Lastname    string           `json:"Lastname"`
	Photo       string           `json:"photo"`
	FrInfoLocal *FriendInfoLocal `json:"fr_info_local"`
}

type SlavesListInfo struct {
	JobName       string
	FetterTime    time.Time `json:"fetter_time"`
	SlaveLevel    int32
	DefenderLevel int32
	FetterType    string
}

type SlaveBuyUpdateInfo struct {
	SlaveId         int32
	JobName         string
	UserType        string
	PurchasePriceSm int64
}
