package domain

// import "time"

// type UserType struct {
// 	name string
// 	id   int32 `pg:"pk"`
// }

// type Fetter struct {
// 	id       int    `pg:"pk"`
// 	ame      string `pg:"type:varchar(255),unique"`
// 	price    int
// 	time     time.Time `pg:"type:timestamp"`
// 	cooldown time.Time `pg:"type:timestamp"`
// }

// type SlaveLevel struct {
// 	id              int `pg:"pk"`
// 	lvl             int `pg:"unique"`
// 	profit          int
// 	money_to_update int `pg:"type:bigint"`
// }

// type SlaveStats struct {
// 	id             int         `pg:"pk"`
// 	level          *SlaveLevel `pg:"fk:id"`
// 	money_quantity int         `pg:"type:bigint"`
// }

// type DefenderLevel struct {
// 	id               int `pg:"pk"`
// 	lvl              int `pg:"unique"`
// 	hp               int
// 	damage           int
// 	damage_to_update int `pg:"type:bigint"`
// }

// type DefenderStats struct {
// 	id              int            `pg:"pk"`
// 	level           *DefenderLevel `pg:"fk:id"`
// 	damage_quantity int            `pg:"type:bigint"`
// }

// type User struct {
// 	id                    int `pg:"pk"`
// 	balance               int `pg:"type:bigint"`
// 	slaves_count          int
// 	income                int            `pg:"type:bigint"`
// 	last_update           time.Time      `pg:"type:timestamp"`
// 	job_name              string         `pg:"type:varchar(255)"`
// 	user_type             *UserType      `pg:"fk:id"`
// 	slave_stats           *SlaveStats    `pg:"rel:has-one"`
// 	defender_stats        *DefenderStats `pg:"rel:has-one"`
// 	purchase_price_silver int            `pg:"type:bigint"`
// 	sale_price_silver     int            `pg:"type:bigint"`
// 	purchase_price_gold   int            `pg:"type:int"`
// 	sale_price_gold       int            `pg:"type:int"`
// 	fetter_time           time.Time      `pg:"type:timestamp"`
// 	fetter_type           *Fetter        `pg:"fk:id"`
// 	has_fetter            bool
// 	slaves                []User `pg:"-"`
// }

// type Slaves struct {
// 	user_id   *User `pg:"fk:id,unique"`
// 	master_id *User `pg:"fk:id"`
// }
