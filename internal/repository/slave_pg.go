package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/jackc/pgx/v4"
)

type SlavePostgres struct {
	db *pgx.Conn
}

func NewSlavePostgres(db *pgx.Conn) *SlavePostgres {
	return &SlavePostgres{db: db}
}

func (rep *SlavePostgres) CreateSlave(slave domain.Slave) error {
	_, err := rep.db.Exec(context.Background(), `INSERT INTO slave(
		user_id, 
		master_id 
	) VALUES ($1, $2);`,
		slave.UserId,
		slave.MasterId)

	return err
}

func (rep *SlavePostgres) GetMaster(userId int32) (int32, error) {
	var masterId int32
	err := rep.db.QueryRow(context.Background(),
		"SELECT * FROM slave WHERE user_id = $1",
		userId).Scan(&masterId)

	return masterId, err
}

func (rep *SlavePostgres) GetSlaves(userId int32) ([]domain.SlavesListInfo, error) {
	slaves := make([]domain.SlavesListInfo, 0, 500)
	rows, err := rep.db.Query(context.Background(),
		`SELECT 
			u.job_name, 
			u.has_fetter,
			sl.lvl,
			dl.lvl,
			sl.profit,
			f.name
		FROM users u
		INNER JOIN slave_stats ss
			ON ss.id = u.slave_stats 
		INNER JOIN slave_level sl
			ON sl.id = ss.level 
		INNER JOIN defender_stats ds
			ON ss.id = u.defender_stats
		INNER JOIN defender_level dl
			ON dl.id = ds.level
		INNER JOIN fetter f
			ON f.id = u.fetter_type;`)

	if err != nil {
		return slaves, err
	}

	defer rows.Close()

	sl := domain.SlavesListInfo{}

	for rows.Next() {
		err := rows.Scan(&sl.JobName,
			&sl.HasFetter,
			&sl.SlaveLevel,
			&sl.SlaveLevel,
			&sl.DefenderLevel,
			&sl.Profit,
			&sl.FetterType)
		if err != nil {
			return slaves, err
		}
		slaves = append(slaves, sl)
	}

	return slaves, nil
}
