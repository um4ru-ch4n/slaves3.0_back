package repository

import (
	"context"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (rep *AuthPostgres) CreateUser(userId int32, userType, fio, photo string) (domain.User, error) {
	user := domain.User{}

	var fetterId int32

	err := rep.db.QueryRow(context.Background(),
		`INSERT INTO users(
			id, 
			user_type,
			fio, 
			photo) 
		VALUES ($1, (SELECT id FROM user_type WHERE name = $2), $3, $4) 
		RETURNING 
			id, 
			fio, 
			photo, 
			balance, 
			gold, 
			last_update, 
			job_name, 
			(SELECT name FROM user_type WHERE id = user_type), 
			slave_level, 
			money_quantity, 
			defender_level, 
			damage_quantity, 
			fetter_time, 
			fetter_type;`,
		userId,
		userType,
		fio,
		photo).Scan(
		&user.Id,
		&user.Fio,
		&user.Photo,
		&user.Balance,
		&user.Gold,
		&user.LastUpdate,
		&user.JobName,
		&user.UserType,
		&user.SlaveLevel,
		&user.MoneyQuantity,
		&user.DefenderLevel,
		&user.DamageQuantity,
		&user.FetterTime,
		&fetterId)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "CreateUser queryRow AuthPostgres")
	}

	err = rep.db.QueryRow(context.Background(),
		`SELECT 
			name, 
			price, 
			duration 
		FROM fetter 
		WHERE id = $1 
		LIMIT 1;`, fetterId).Scan(
		&user.FetterType,
		&user.FetterPrice,
		&user.FetterDuration)

	return user, errors.Wrap(err, "select fetter queryRow AuthPostgres")
}

func (rep *AuthPostgres) GetUser(id int32) (domain.User, error) {
	user := domain.User{}
	err := rep.db.QueryRow(context.Background(),
		`SELECT 
			u.id, 
			u.fio, 
			u.photo, 
			u.balance, 
			u.gold, 
			u.last_update, 
			u.job_name, 
			ut.name, 
			u.slave_level, 
			u.money_quantity, 
			u.defender_level, 
			u.damage_quantity, 
			u.fetter_time, 
			f.name, 
			f.price, 
			f.duration, 
			CASE WHEN um.master_id is NULL THEN 0 ELSE um.master_id END AS master_id, 
			CASE WHEN master_id is NULL THEN '' ELSE (SELECT us.fio FROM users us WHERE us.id = um.master_id) END AS master_fio 
		FROM users u 
		INNER JOIN user_type ut 
			ON ut.id = u.user_type 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type 
		LEFT JOIN user_master um 
			ON um.user_id = u.id 
		WHERE u.id = $1 
		LIMIT 1;`, id).Scan(
		&user.Id,
		&user.Fio,
		&user.Photo,
		&user.Balance,
		&user.Gold,
		&user.LastUpdate,
		&user.JobName,
		&user.UserType,
		&user.SlaveLevel,
		&user.MoneyQuantity,
		&user.DefenderLevel,
		&user.DamageQuantity,
		&user.FetterTime,
		&user.FetterType,
		&user.FetterPrice,
		&user.FetterDuration,
		&user.MasterId,
		&user.MasterFio)

	return user, errors.Wrap(err, "GetUser queryRow AuthPostgres")
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

	return usType, errors.Wrap(err, "GetUserType queryRow AuthPostgres")
}

func (rep *AuthPostgres) GetFriendsInfo(ids []int) (map[int32]domain.FriendInfo, error) {
	friendsInfo := make(map[int32]domain.FriendInfo)

	rows, err := rep.db.Query(context.Background(),
		`SELECT 
			u.id, 
			u.fio, 
			u.photo, 
			CASE WHEN um.master_id is NULL THEN 0 ELSE um.master_id END AS master_id, 
			u.fetter_time, 
			f.name, 
			f.duration, 
			u.slave_level, 
			u.defender_level 
		FROM users u 
		LEFT JOIN user_master um 
			ON um.user_id = u.id 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type 
		WHERE u.id = ANY ($1);`, ids)

	if err != nil {
		return friendsInfo, errors.Wrap(err, "GetFriendsInfo query AuthPostgres")
	}

	defer rows.Close()

	fr := domain.FriendInfo{}

	for rows.Next() {
		err := rows.Scan(
			&fr.Id,
			&fr.Fio,
			&fr.Photo,
			&fr.MasterId,
			&fr.FetterTime,
			&fr.FetterType,
			&fr.FetterDuration,
			&fr.SlaveLevel,
			&fr.DefenderLevel)

		if err != nil {
			return friendsInfo, errors.Wrap(err, "GetFriendsInfo scan AuthPostgres")
		}

		friendsInfo[fr.Id] = fr
	}

	return friendsInfo, nil
}

func (rep *AuthPostgres) SlaveBuyUpdateInfo(newData domain.SlaveBuyUpdateInfo) error {
	_, err := rep.db.Exec(context.Background(),
		`UPDATE users 
		SET 
			job_name = $1, 
			user_type = (SELECT id FROM user_type WHERE name = $2)
		WHERE id = $3;`,
		newData.JobName,
		newData.UserType,
		newData.SlaveId)

	return errors.Wrap(err, "SlaveBuyUpdateInfo exec AuthPostgres")
}

func (rep *AuthPostgres) UserBalanceUpdate(userId int32, balance int64, gold int32) error {
	_, err := rep.db.Exec(context.Background(),
		`UPDATE users 
		SET 
			balance = $1, 
			gold = $2 
		WHERE id = $3;`,
		balance,
		gold,
		userId)

	return errors.Wrap(err, "UserBalanceUpdate exec AuthPostgres")
}

func (rep *AuthPostgres) SetFetterTime(userId int32, fetterTime time.Time) error {
	_, err := rep.db.Exec(context.Background(),
		"UPDATE users SET fetter_time = $1 WHERE id = $2", fetterTime, userId)

	return errors.Wrap(err, "SetFetterTime exec AuthPostgres")
}

func (rep *AuthPostgres) GetFetterTime(userId int32) (time.Time, error) {
	var fetterTime time.Time
	err := rep.db.QueryRow(context.Background(),
		`SELECT fetter_time FROM users WHERE id = $1 LIMIT 1;`,
		userId).Scan(&fetterTime)

	return fetterTime, errors.Wrap(err, "SetFetterTime queryRow AuthPostgres")
}

func (rep *AuthPostgres) GetUserBalance(userId int32) (int64, int32, error) {
	var gold int32
	var balance int64
	err := rep.db.QueryRow(context.Background(), `SELECT balance, gold FROM users WHERE id = $1 LIMIT 1;`,
		userId).Scan(&balance, &gold)

	return balance, gold, errors.Wrap(err, "GetUserBalance queryRow AuthPostgres")
}

func (rep *AuthPostgres) GetRatingBySlavesCount(limit int32) ([]domain.RatingSlavesCount, error) {
	ratingSlavesCount := make([]domain.RatingSlavesCount, 0, limit)

	rows, err := rep.db.Query(context.Background(),
		`SELECT 
			um.master_id,  
			u.fio, 
			count(um.master_id) as slaves_count, 
			u.photo, 
			u.fetter_time, 
			f.name, 
			f.duration 
		FROM user_master um 
		INNER JOIN users u
			ON u.id = um.master_id 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type  
		GROUP BY 
			um.master_id, 
			u.fio, 
			u.photo, 
			u.fetter_time, 
			f.name, 
			f.duration  
		ORDER BY slaves_count DESC 
		LIMIT $1 ;`, limit)

	if err != nil {
		return ratingSlavesCount, errors.Wrap(err, "GetRatingBySlavesCount query AuthPostgres")
	}

	defer rows.Close()

	rs := domain.RatingSlavesCount{}

	for rows.Next() {
		err := rows.Scan(
			&rs.Id,
			&rs.Fio,
			&rs.SlavesCount,
			&rs.Photo,
			&rs.FetterTime,
			&rs.FetterType,
			&rs.FetterDuration)

		if err != nil {
			return ratingSlavesCount, errors.Wrap(err, "GetUserBalance rows.Scan AuthPostgres")
		}

		ratingSlavesCount = append(ratingSlavesCount, rs)
	}

	return ratingSlavesCount, nil
}

func (rep *AuthPostgres) SetJobName(slaveId int32, jobName string) error {
	_, err := rep.db.Exec(context.Background(),
		"UPDATE users SET job_name = $1 WHERE id = $2;", jobName, slaveId)

	return errors.Wrap(err, "SetJobName exec AuthPostgres")
}

func (rep *AuthPostgres) GetLastUpdate(userId int32) (time.Time, error) {
	var lastUpdate time.Time

	err := rep.db.QueryRow(context.Background(),
		"SELECT last_update FROM users WHERE id = $1 LIMIT 1;", userId).Scan(&lastUpdate)

	return lastUpdate, errors.Wrap(err, "GetLastUpdate queryRow AuthPostgres")
}

func (rep *AuthPostgres) UpdateUserBalanceHour(userId int32, balance int64) error {
	_, err := rep.db.Exec(context.Background(),
		"UPDATE users SET balance = $1, last_update = NOW() WHERE id = $2;", balance, userId)

	return errors.Wrap(err, "UpdateUserBalanceHour exec AuthPostgres")
}

func (rep *AuthPostgres) UpdateSlaveHour(slaveInfo domain.SlaveInfoForUpdate) error {
	_, err := rep.db.Exec(context.Background(),
		"UPDATE users SET slave_level = $1, money_quantity = $2 WHERE id = $3;",
		slaveInfo.SlaveLevel, slaveInfo.MoneyQuantity, slaveInfo.Id)

	return errors.Wrap(err, "UpdateSlaveHour exec AuthPostgres")
}
