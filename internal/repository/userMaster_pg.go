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
			u.job_name, 
			u.fetter_time,
			u.slave_level,
			u.defender_level,
			f.name 
		FROM users u 
		INNER JOIN fetter f 
			ON f.id = u.fetter_type;`)

	if err != nil {
		return slaves, err
	}

	defer rows.Close()

	sl := domain.SlavesListInfo{}

	for rows.Next() {
		err := rows.Scan(&sl.JobName,
			&sl.FetterTime,
			&sl.SlaveLevel,
			&sl.DefenderLevel,
			&sl.FetterType)
		if err != nil {
			return slaves, err
		}
		slaves = append(slaves, sl)
	}

	return slaves, nil
}
