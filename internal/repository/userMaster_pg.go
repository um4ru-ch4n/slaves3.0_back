package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type UserMasterPostgres struct {
	db *pgx.Conn
}

func NewUserMasterPostgres(db *pgx.Conn) *UserMasterPostgres {
	return &UserMasterPostgres{db: db}
}

func (rep *UserMasterPostgres) CreateOrUpdateSlave(userId int32, masterId int32) error {
	_, err := rep.db.Exec(context.Background(),
		`INSERT INTO user_master(
			user_id, 
			master_id 
		) VALUES ($1, $2)
		ON CONFLICT (user_id)
		DO
			UPDATE SET master_id = $2;`,
		userId,
		masterId)

	return err
}

func (rep *UserMasterPostgres) GetMaster(userId int32) (int32, error) {
	var masterId int32
	err := rep.db.QueryRow(context.Background(),
		"SELECT master_id FROM user_master WHERE user_id = $1 LIMIT 1;",
		userId).Scan(&masterId)

	return masterId, err
}

func (rep *UserMasterPostgres) GetSlaves(userId int32) ([]domain.SlavesListInfo, error) {
	slaves := make([]domain.SlavesListInfo, 0, 500)
	rows, err := rep.db.Query(context.Background(),
		`SELECT 
			u.id, 
			u.fio, 
			u.photo, 
			u.job_name, 
			u.fetter_time,
			f.name, 
			f.duration,  
			u.slave_level,
			u.defender_level 
		FROM user_master um 
		INNER JOIN users u 
			ON u.id = um.user_id 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type
		WHERE um.master_id = $1;`, userId)

	if err != nil {
		return slaves, err
	}

	defer rows.Close()

	sl := domain.SlavesListInfo{}

	for rows.Next() {
		err := rows.Scan(
			&sl.Id,
			&sl.Fio,
			&sl.Photo,
			&sl.JobName,
			&sl.FetterTime,
			&sl.FetterType,
			&sl.FetterDuration,
			&sl.SlaveLevel,
			&sl.DefenderLevel)
		if err != nil {
			return slaves, err
		}
		slaves = append(slaves, sl)
	}

	return slaves, nil
}

func (rep *UserMasterPostgres) SaleSlave(slaveId int32) error {
	_, err := rep.db.Exec(context.Background(), "DELETE FROM user_master WHERE user_id = $1;", slaveId)

	return err
}

func (rep *UserMasterPostgres) GetSlavesForUpdate(userId int32) ([]domain.SlaveInfoForUpdate, error) {
	slaves := make([]domain.SlaveInfoForUpdate, 0, 500)
	rows, err := rep.db.Query(context.Background(),
		`SELECT 
			u.id, 
			(SELECT ut.name FROM user_type ut WHERE ut.id = u.user_type) as user_type, 
			u.slave_level, 
			u.money_quantity 
		FROM user_master um 
		INNER JOIN users u 
			ON u.id = um.user_id 
		WHERE um.master_id = $1;`, userId)

	if err != nil {
		return slaves, err
	}

	defer rows.Close()

	sl := domain.SlaveInfoForUpdate{}

	for rows.Next() {
		err := rows.Scan(
			&sl.Id,
			&sl.UserType,
			&sl.SlaveLevel,
			&sl.MoneyQuantity)
		if err != nil {
			return slaves, err
		}

		slaves = append(slaves, sl)
	}

	return slaves, nil
}
