package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(connString string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 20
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

type DataBase struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DataBase {
	return &DataBase{pool: pool}
}
