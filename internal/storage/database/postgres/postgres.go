package postgres

import (
	"context"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	db *pgxpool.Pool
}

func NewDB(addrConn string) (*DB, error) {
	addr, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		logger.Errorf("error parse config: %s", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), addr)
	if err != nil {
		logger``.Errorf("Error create NewWithConfig : %s", err)
	}

	db := &DB{
		db: conn,
	}

	return db, nil
}

func (db *DB) Save(saveStr string) (string, error) {
	return saveStr, nil //TODO improve in future
}

func (db *DB) Get(getStr string) string {
	return getStr //TODO improve in future
}

func (db *DB) Close() error {
	return nil //TODO improve in future
}
