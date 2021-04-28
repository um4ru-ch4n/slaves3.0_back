package repository

import (
	"context"

	"github.com/00mrx00/slaves3.0_back/internal/config"
	"github.com/jackc/pgx/v4"
)

func NewPostgresDB(cfg config.DbConfig) (*pgx.Conn, error) {
	baseconn := "host=" + cfg.Host + " port=" + cfg.Port + " user=" + cfg.Username +
		" password=" + cfg.Password + " database=" + cfg.DbName + " sslmode=" + cfg.SSLMode

	db, err := pgx.Connect(context.Background(), baseconn)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateSchema(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `
	CREATE TABLE public.user_type
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
    cooldown integer NOT NULL,
    CONSTRAINT fetter_pkey PRIMARY KEY (id),
    CONSTRAINT fetter_name_unique UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.fetter
    OWNER to postgres;

	CREATE TABLE public.slave_level
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    lvl integer NOT NULL,
    profit integer NOT NULL,
    money_to_update bigint NOT NULL,
    CONSTRAINT slave_level_pkey PRIMARY KEY (id),
    CONSTRAINT slave_level_lvl_unique UNIQUE (lvl)
)

TABLESPACE pg_default;

ALTER TABLE public.slave_level
    OWNER to postgres;

	CREATE TABLE public.slave_stats
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    level integer NOT NULL,
    money_quantity bigint NOT NULL,
    CONSTRAINT slave_stats_pkey PRIMARY KEY (id),
    CONSTRAINT slave_stats_level_fk FOREIGN KEY (level)
        REFERENCES public.slave_level (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.slave_stats
    OWNER to postgres;

	CREATE TABLE public.defender_level
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    lvl integer NOT NULL,
    hp integer NOT NULL,
    damage integer NOT NULL,
    damage_to_update bigint NOT NULL,
    CONSTRAINT defender_level_pkey PRIMARY KEY (id),
    CONSTRAINT defender_level_lvl_unique UNIQUE (lvl)
)

TABLESPACE pg_default;

ALTER TABLE public.defender_level
    OWNER to postgres;

	CREATE TABLE public.defender_stats
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    level integer NOT NULL,
    damage_quantity bigint NOT NULL,
    CONSTRAINT defender_stats_pkey PRIMARY KEY (id),
    CONSTRAINT defender_stats_level_fk FOREIGN KEY (level)
        REFERENCES public.defender_level (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.defender_stats
    OWNER to postgres;

	CREATE TABLE public.users
(
    id integer NOT NULL,
    slaves_count integer NOT NULL,
    balance bigint NOT NULL,
    income bigint NOT NULL,
    last_update timestamp without time zone NOT NULL,
    job_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_type integer NOT NULL,
    slave_stats integer NOT NULL,
    defender_stats integer NOT NULL,
    purchase_price_sm bigint NOT NULL,
    sale_price_sm bigint NOT NULL,
    purchase_price_gm integer NOT NULL,
    sale_price_gm integer NOT NULL,
    has_fetter boolean NOT NULL,
    fetter_time timestamp without time zone NOT NULL,
    fetter_type integer NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_defender_stats_fk FOREIGN KEY (defender_stats)
        REFERENCES public.defender_stats (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION,
    CONSTRAINT user_fetter_type_fk FOREIGN KEY (fetter_type)
        REFERENCES public.fetter (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION,
    CONSTRAINT user_slave_stats_fk FOREIGN KEY (slave_stats)
        REFERENCES public.slave_stats (id) MATCH SIMPLE
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

	CREATE TABLE public.slave
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

ALTER TABLE public.slave
    OWNER to postgres;
	`)
	return err
}
