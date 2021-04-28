package repository

import (
	"context"
	"fmt"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (rep *AuthPostgres) CreateUser(user domain.User) error {
	_, err := rep.db.Exec(context.Background(), `INSERT INTO users(
		id, 
		slaves_count, 
		balance, 
		income,
		last_update,
		job_name,
		user_type,
		slave_stats,
		defender_stats,
		purchase_price_sm,
		sale_price_sm,
		purchase_price_gm,
		sale_price_gm, 
		has_fetter, 
		fetter_time, 
		fetter_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);`,
		user.Id,
		user.SlavesCount,
		user.Balance,
		user.Income,
		user.LastUpdate,
		user.JobName,
		user.UserType.Id,
		user.SlaveStats.Id,
		user.DefenderStats.Id,
		user.PurchasePriceSm,
		user.SalePriceSm,
		user.PurchasePriceGm,
		user.SalePriceGm,
		user.HasFetter,
		user.FetterTime,
		user.FetterType.Id)

	return err
}

func (rep *AuthPostgres) GetUser(id int32) (domain.User, error) {
	user := domain.User{
		UserType: &domain.UserType{},
		SlaveStats: &domain.SlaveStats{
			Level: &domain.SlaveLevel{},
		},
		DefenderStats: &domain.DefenderStats{
			Level: &domain.DefenderLevel{},
		},
		FetterType: &domain.Fetter{},
	}
	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			u.id,
			u.slaves_count,
			u.balance,
			u.income,
			u.last_update,
			u.job_name,
			u.purchase_price_sm,
			u.sale_price_sm,
			u.purchase_price_gm,
			u.sale_price_gm,
			u.has_fetter,
			u.fetter_time,
			ut.id,
			ut.name,
			ss.id,
			ss.money_quantity,
			sl.id,
			sl.lvl,
			sl.profit,
			sl.money_to_update,
			ds.id,
			ds.damage_quantity,
			dl.id,
			dl.lvl,
			dl.hp,
			dl.damage,
			dl.damage_to_update,
			f.id,
			f.name,
			f.price,
			f.duration,
			f.cooldown 
		FROM users u 
		INNER JOIN user_type ut 
			ON ut.id = u.user_type 
		INNER JOIN slave_stats ss 
			ON ss.id = u.slave_stats 
		INNER JOIN slave_level sl 
			ON sl.id = ss.level 
		INNER JOIN defender_stats ds 
			ON ds.id = u.defender_stats 
		INNER JOIN defender_level dl 
			ON dl.id = ds.level 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type 
		WHERE u.id = $1;`, id).Scan(
		&user.Id,
		&user.SlavesCount,
		&user.Balance,
		&user.Income,
		&user.LastUpdate,
		&user.JobName,
		&user.PurchasePriceSm,
		&user.SalePriceSm,
		&user.PurchasePriceGm,
		&user.SalePriceGm,
		&user.HasFetter,
		&user.FetterTime,
		&user.UserType.Id,
		&user.UserType.Name,
		&user.SlaveStats.Id,
		&user.SlaveStats.MoneyQuantity,
		&user.SlaveStats.Level.Id,
		&user.SlaveStats.Level.Lvl,
		&user.SlaveStats.Level.Profit,
		&user.SlaveStats.Level.MoneyToUpdate,
		&user.DefenderStats.Id,
		&user.DefenderStats.DamageQuantity,
		&user.DefenderStats.Level.Id,
		&user.DefenderStats.Level.Lvl,
		&user.DefenderStats.Level.Hp,
		&user.DefenderStats.Level.Damage,
		&user.DefenderStats.Level.DamageToUpdate,
		&user.FetterType.Id,
		&user.FetterType.Name,
		&user.FetterType.Price,
		&user.FetterType.Duration,
		&user.FetterType.Cooldown)

	fmt.Println(err)

	return user, err
}
