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
	Cooldown int32  `json:"cooldown"`
}

type SlaveLevel struct {
	Id            int32 `json:"id"`
	Lvl           int32 `json:"lvl"`
	Profit        int32 `json:"profit"`
	MoneyToUpdate int64 `json:"money_to_update"`
}

type SlaveStats struct {
	Id            int32       `json:"id"`
	Level         *SlaveLevel `json:"level"`
	MoneyQuantity int64       `json:"money_quantity"`
}

type DefenderLevel struct {
	Id             int32 `json:"id"`
	Lvl            int32 `json:"lvl"`
	Hp             int32 `json:"hp"`
	Damage         int32 `json:"damage"`
	DamageToUpdate int64 `json:"damage_to_update"`
}

type DefenderStats struct {
	Id             int32          `json:"id"`
	Level          *DefenderLevel `json:"level"`
	DamageQuantity int64          `json:"damage_quantity"`
}

type User struct {
	Id              int32          `json:"id"`
	SlavesCount     int32          `json:"slaves_count"`
	Balance         int64          `json:"balance"`
	Income          int64          `json:"income"`
	LastUpdate      time.Time      `json:"last_update"`
	JobName         string         `json:"job_name"`
	UserType        *UserType      `json:"user_type"`
	SlaveStats      *SlaveStats    `json:"slave_stats"`
	DefenderStats   *DefenderStats `json:"defender_stats"`
	PurchasePriceSm int64          `json:"purchase_price_sm"`
	SalePriceSm     int64          `json:"sale_price_sm"`
	PurchasePriceGm int32          `json:"purchase_price_gm"`
	SalePriceGm     int32          `json:"sale_price_gm"`
	HasFetter       bool           `json:"has_fetter"`
	FetterTime      time.Time      `json:"fetter_time"`
	FetterType      *Fetter        `json:"fetter_type"`
	VkInfo          *UserVkInfo    `json:"vk_info"`
}

type Slave struct {
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
	MasterId        int32  `json:"master_id"`
	MasterFirstname string `json:"master_Firstname"`
	MasterLastname  string `json:"master_Lastname"`
	HasFetter       bool   `json:"has_fetter"`
	FetterType      string `json:"fetter_type"`
	PurchasePriceSm int64  `json:"purchase_price_sm"`
	PurchasePriceGm int32  `json:"purchase_price_gm"`
	SlaveLevel      int32  `json:"slave_level"`
	DefenderLevel   int32  `json:"defender_level"`
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
	HasFetter     bool
	SlaveLevel    int32
	DefenderLevel int32
	Profit        int32
	FetterType    string
}
