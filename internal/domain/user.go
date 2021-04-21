package domain

import (
	"time"

	_ "github.com/go-pg/pg/v10"
)

type UserType struct {
	Id   int32 `pg:"id,pk"`
	Name string
}

type Fetter struct {
	Id       int32  `pg:"id,pk"`
	Name     string `pg:"type:varchar(255),unique"`
	Price    int32
	Time     time.Time `pg:"type:timestamp"`
	Cooldown time.Time `pg:"type:timestamp"`
}

type SlaveLevel struct {
	Id              int32 `pg:"id,pk"`
	Lvl             int32 `pg:",unique"`
	Profit          int32
	Money_to_update int64 `pg:"type:bigint"`
}

type SlaveStats struct {
	Id             int32       `pg:"id,pk"`
	Level          *SlaveLevel `pg:"rel:has-one"`
	Money_quantity int64       `pg:"type:bigint"`
}

type DefenderLevel struct {
	Id               int32 `pg:"id,pk"`
	Lvl              int32 `pg:",unique"`
	Hp               int32
	Damage           int32
	Damage_to_update int64 `pg:"type:bigint"`
}

type DefenderStats struct {
	Id              int32          `pg:"id,pk"`
	Level           *DefenderLevel `pg:"rel:has-one"`
	Damage_quantity int64          `pg:"type:bigint"`
}

type User struct {
	Id                    int32 `pg:"id,pk"`
	Slaves_count          int32
	Balance               int64          `pg:"type:bigint"`
	Income                int64          `pg:"type:bigint"`
	Last_update           time.Time      `pg:"type:timestamp"`
	Job_name              string         `pg:"type:varchar(255)"`
	User_type             *UserType      `pg:"rel:has-one"`
	Slave_stats           *SlaveStats    `pg:"rel:has-one"`
	Defender_stats        *DefenderStats `pg:"rel:has-one"`
	Purchase_price_silver int64          `pg:"type:bigint"`
	Sale_price_silver     int64          `pg:"type:bigint"`
	Purchase_price_gold   int32
	Sale_price_gold       int32
	Has_fetter            bool
	Fetter_time           time.Time `pg:"type:timestamp"`
	Fetter_type           *Fetter   `pg:"rel:has-one"`
	Slaves                []User    `pg:"-"`
}

type Slaves struct {
	User_id   *User `pg:"fk:id,unique"`
	Master_id *User `pg:"fk:id"`
}
