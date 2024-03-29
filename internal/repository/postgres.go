package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func NewPostgresDB(cfg config.DbConfig) (*pgxpool.Pool, error) {
	baseconn := "host=" + cfg.Host + " port=" + cfg.Port + " user=" + cfg.Username +
		" password=" + cfg.Password + " database=" + cfg.DbName + " sslmode=" + cfg.SSLMode

	db, err := pgxpool.Connect(context.Background(), baseconn)

	if err != nil {
		return nil, errors.Wrap(err, "db connect failed")
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "db ping failed")
	}

	return db, nil
}

func CreateSchema(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(),
		`CREATE TABLE public.user_type
        (
            id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
            name character varying(255) COLLATE pg_catalog."default" NOT NULL,
            CONSTRAINT user_type_pkey PRIMARY KEY (id),
            CONSTRAINT user_type_name_unique UNIQUE (name)
        )
    
        TABLESPACE pg_default;
    
        ALTER TABLE public.user_type
            OWNER to postgres;
    
        	CREATE TABLE public.fetter
        (
            id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
            name character varying(255) COLLATE pg_catalog."default" NOT NULL,
            price integer NOT NULL,
            duration integer NOT NULL,
            CONSTRAINT fetter_pkey PRIMARY KEY (id),
            CONSTRAINT fetter_name_unique UNIQUE (name)
        )
    
        TABLESPACE pg_default;
    
        ALTER TABLE public.fetter
            OWNER to postgres;
    
        	CREATE TABLE public.users
        (
            id integer NOT NULL,
            fio character varying(255) COLLATE pg_catalog."default" NOT NULL,
            photo character varying(255) COLLATE pg_catalog."default" NOT NULL,
            balance bigint NOT NULL DEFAULT 100,
            gold integer NOT NULL DEFAULT 0,
            last_update timestamp with time zone NOT NULL DEFAULT NOW(),
            job_name character varying(255) NOT NULL DEFAULT '',
            user_type integer NOT NULL DEFAULT 1,
            slave_level integer NOT NULL DEFAULT 1,
            money_quantity bigint NOT NULL DEFAULT 0,
            defender_level integer NOT NULL DEFAULT 1,
            damage_quantity bigint NOT NULL DEFAULT 0,
            fetter_time timestamp with time zone NOT NULL DEFAULT '1971-11-03T00:00:00.0000+03:00'::timestamp,
            fetter_type integer NOT NULL DEFAULT 1,
            CONSTRAINT user_pkey PRIMARY KEY (id),
            CONSTRAINT user_fetter_type_fk FOREIGN KEY (fetter_type)
                REFERENCES public.fetter (id) MATCH SIMPLE
                ON UPDATE CASCADE
                ON DELETE NO ACTION,
            CONSTRAINT user_user_type_fk FOREIGN KEY (user_type)
                REFERENCES public.user_type (id) MATCH SIMPLE
                ON UPDATE CASCADE
                ON DELETE NO ACTION
        )
    
        TABLESPACE pg_default;
    
        ALTER TABLE public.users
            OWNER to postgres;
    
        	CREATE TABLE public.user_master
        (
            user_id integer NOT NULL,
            master_id integer NOT NULL,
            CONSTRAINT slave_user_id_unique UNIQUE (user_id),
            CONSTRAINT slave_master_id_fk FOREIGN KEY (master_id)
                REFERENCES public.users (id) MATCH SIMPLE
                ON UPDATE CASCADE
                ON DELETE CASCADE,
            CONSTRAINT slave_user_id_fk FOREIGN KEY (user_id)
                REFERENCES public.users (id) MATCH SIMPLE
                ON UPDATE CASCADE
                ON DELETE CASCADE
        )
    
        TABLESPACE pg_default;
    
        ALTER TABLE public.user_master
            OWNER to postgres;`)

	return errors.Wrap(err, "create initial schema exec failed")
}

func CreateUserTypes(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(),
		`INSERT INTO user_type(name) 
        VALUES  
            ('simp'), 
            ('slave'), 
            ('defender');`)

	return errors.Wrap(err, "create initial userTypes failed")
}

func CreateFetter(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(),
		`INSERT INTO fetter(name, price, duration) 
        VALUES 
            ('common', 80, 120), 
            ('uncommon', 100, 240), 
            ('rare', 120, 360), 
            ('epic', 140, 480), 
            ('immortal', 160, 720), 
            ('legendary', 180, 1440);`)

	return errors.Wrap(err, "create initial fetter failed")
}
