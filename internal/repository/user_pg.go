package repository

import (
	"context"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (rep *AuthPostgres) CreateUser(userId int32, userType string) (domain.User, error) {
	user := domain.User{
		FetterType: &domain.Fetter{},
	}

	err := rep.db.QueryRow(context.Background(),
		`INSERT INTO users(
			id, 
			user_type) 
		VALUES ($1, (SELECT id FROM user_type WHERE name = $2)) 
		RETURNING 
			id, 
			balance, 
			gold, 
			last_update, 
			job_name, 
			(SELECT name FROM user_type WHERE id = user_type), 
			slave_level, 
			money_quantity, 
			defender_level, 
			damage_quantity,  
			purchase_price_sm, 
			purchase_price_gm, 
			fetter_time, 
			fetter_type;`,
		userId,
		userType).Scan(
		&user.Id,
		&user.Balance,
		&user.Gold,
		&user.LastUpdate,
		&user.JobName,
		&user.UserType,
		&user.SlaveLevel,
		&user.MoneyQuantity,
		&user.DefenderLevel,
		&user.DamageQuantity,
		&user.PurchasePriceSm,
		&user.PurchasePriceGm,
		&user.FetterTime,
		&user.FetterType.Id)

	return user, err
}

func (rep *AuthPostgres) GetUser(id int32) (domain.User, error) {
	user := domain.User{
		FetterType: &domain.Fetter{},
	}
	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			u.id, 
			u.balance, 
			u.gold, 
			u.last_update, 
			u.job_name, 
			ut.name, 
			u.slave_level, 
			u.money_quantity, 
			u.defender_level, 
			u.damage_quantity, 
			u.purchase_price_sm, 
			u.purchase_price_gm, 
			u.fetter_time, 
			f.id, 
			f.name, 
			f.price, 
			f.duration 
		FROM users u 
		INNER JOIN user_type ut 
			ON ut.id = u.user_type 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type 
		WHERE u.id = $1 
		LIMIT 1;`, id).Scan(
		&user.Id,
		&user.Balance,
		&user.Gold,
		&user.LastUpdate,
		&user.JobName,
		&user.UserType,
		&user.SlaveLevel,
		&user.MoneyQuantity,
		&user.DefenderLevel,
		&user.DamageQuantity,
		&user.PurchasePriceSm,
		&user.PurchasePriceGm,
		&user.FetterTime,
		&user.FetterType.Id,
		&user.FetterType.Name,
		&user.FetterType.Price,
		&user.FetterType.Duration)

	return user, err
}

func (rep *AuthPostgres) GetUserType(userId int32) (string, error) {
	var usType string
	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			ut.name 
		FROM users u
		INNER JOIN user_type ut 
			ON ut.id = u.user_type
		WHERE id = $1
		LIMIT 1;`, userId).Scan(&usType)

	return usType, err
}

func (rep *AuthPostgres) GetFriendInfoLocal(id int32) (domain.FriendInfoLocal, error) {
	friendInfoLocal := domain.FriendInfoLocal{}

	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			um.master_id, 
			u.fetter_time, 
			f.name, 
			u.purchase_price_sm, 
			u.purchase_price_gm, 
			u.slave_level, 
			u.defender_level 
		FROM users u 
		INNER JOIN user_master um 
			ON um.user_id = u.id 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type 
		WHERE u.id = $1 
		LIMIT 1;`, id).Scan(
		&friendInfoLocal.MasterId,
		&friendInfoLocal.FetterTime,
		&friendInfoLocal.FetterType,
		&friendInfoLocal.PurchasePriceSm,
		&friendInfoLocal.PurchasePriceGm,
		&friendInfoLocal.SlaveLevel,
		&friendInfoLocal.DefenderLevel,
	)

	return friendInfoLocal, err
}

func (rep *AuthPostgres) SlaveBuyUpdateInfo(newData domain.SlaveBuyUpdateInfo) error {
	_, err := rep.db.Exec(context.Background(),
		`UPDATE users 
		SET 
			job_name = $1, 
			user_type = (SELECT id FROM user_type WHERE name = $2), 
			purchase_price_sm = $3 
		WHERE id = $4;`,
		newData.JobName,
		newData.UserType,
		newData.PurchasePriceSm,
		newData.SlaveId)

	return err
}

func (rep *AuthPostgres) SlaveBalanceUpdate(userId int32, balance int64, gold int32) error {
	_, err := rep.db.Exec(context.Background(),
		`UPDATE users 
		SET 
			balance = $1, 
			gold = $2 
		WHERE id = $3;`,
		balance,
		gold,
		userId)

	return err
}

func (rep *AuthPostgres) SetFetterTime(userId int32, fetterTime time.Time) error {
	_, err := rep.db.Exec(context.Background(),
		"UPDATE users SET fetter_time = $1 WHERE id = $2", fetterTime, userId)

	return err
}

func (rep *AuthPostgres) GetFetterTime(userId int32) (time.Time, error) {
	var fetterTime time.Time
	err := rep.db.QueryRow(context.Background(),
		`SELECT fetter_time FROM users WHERE id = $1 LIMIT 1;`,
		userId).Scan(&fetterTime)

	return fetterTime, err
}

func (rep *AuthPostgres) GetUserBalance(userId int32) (int64, int32, error) {
	var gold int32
	var balance int64
	err := rep.db.QueryRow(context.Background(), `SELECT balance, gold FROM users WHERE id = $1 LIMIT 1;`,
		userId).Scan(&balance, &gold)

	return balance, gold, err
}
