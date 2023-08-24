package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"github.com/kisanetik/learn_go_inc1/internal/utils"
	uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

type DB struct {
	db *pgxpool.Pool
}

func NewPostgresDB(addrConn string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		return nil, fmt.Errorf("error parse config: %w", err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	db, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("error new config: %w", err)
	}

	psql := &DB{db: db}

	exists, err := psql.checkIsTablesExists()
	if err != nil {
		return nil, fmt.Errorf("error check is table exists: %w", err)
	}

	if !exists {
		err = psql.createTable()
		if err != nil {
			return nil, fmt.Errorf("error create table: %w", err)
		}
	}

	return psql, nil
}

func (psql *DB) Save(longURL, corrID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var count string

	shortURL := utils.RandomString()

	s := psql.db.QueryRow(ctx, `SELECT COUNT(*) FROM yandex`)

	err := s.Scan(&count)
	if err != nil {
		logger.Errorf("error in Scan count in SELECT query: %s", err)
	}

	if corrID == "" {
		corrID = shortURL
	}

	_, err = psql.db.Exec(ctx, `INSERT INTO yandex (id, longurl, shorturl, correlation) VALUES ($1, $2, $3, $4);`, count, longURL, shortURL, corrID)
	if err != nil {
		return "", fmt.Errorf("error is INSERT data in database: %w", err)
	}

	return shortURL, nil
}

func (psql *DB) Get(shortURL, corrID string) (string, string) {
	var longURL string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	row := psql.db.QueryRow(ctx, `SELECT longurl FROM yandex WHERE shorturl = $1`, shortURL)

	err := row.Scan(&longURL)
	if err != nil {
		logger.Errorf("error in Scan longURL in SELECT query: %s", err)
	}

	return longURL, corrID
}

func (psql *DB) Close() error {
	psql.db.Close()
	return nil
}

func (psql *DB) createTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := psql.db.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS yandex (
    		id SERIAL PRIMARY KEY,
   			longurl VARCHAR(255) NOT NULL,
    		shorturl VARCHAR(255) NOT NULL,
   			correlation VARCHAR(255) NOT NULL);`)

	return err
}

func (psql *DB) checkIsTablesExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	row := psql.db.QueryRow(ctx, `SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'yandex')`)

	var res bool

	err := row.Scan(&res)
	if err != nil {
		return false, fmt.Errorf("error scan: %w", err)
	}

	return res, nil
}

func (psql *DB) CheckIsURLExists(longURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	row := psql.db.QueryRow(ctx, `SELECT shorturl FROM yandex WHERE longurl = $1`, longURL)

	var res string

	err := row.Scan(&res)
	if err != nil {
		return "", fmt.Errorf("error in Scan res in SELECT query: %w", err)
	}

	return res, nil
}

func (psql *DB) Ping() bool {
	return psql.db.Ping(context.Background()) == nil
}
