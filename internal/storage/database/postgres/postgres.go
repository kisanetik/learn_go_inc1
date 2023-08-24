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

var Count uint64

func NewPostgresDB(addrConn string) (*DB, error) {
	conn, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		return nil, fmt.Errorf("error parse config: %w", err)
	}

	conn.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	db, err := pgxpool.NewWithConfig(context.Background(), conn)
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

	Count++

	return psql, nil
}

func (psql *DB) Save(longURL string) (string, error) {
	shortURL := utils.RandomString()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := psql.db.Exec(ctx, `INSERT INTO yandex (id, longurl, shorturl) VALUES ($1, $2, $3);`, Count, longURL, shortURL)
	if err != nil {
		return "", fmt.Errorf("error is INSERT data in database: %w", err)
	}

	Count++
	return shortURL, nil
}

func (psql *DB) Get(shortURL string) string {
	var longURL string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	row := psql.db.QueryRow(ctx, `SELECT longurl FROM yandex WHERE shorturl = $1`, shortURL)

	err := row.Scan(&longURL)
	if err != nil {
		logger.Errorf("error is Scan data in SELECT Query: %s", err)
		return ""
	}

	return longURL
}

func (psql *DB) Close() error {
	//if err := psql.db.Close; err != nil {
	//	logger.Errorf("error close db: %s", err)
	//}

	return nil //TODO заглушка на будущее, кажется что в бд этот метод вообще не нужен
}
func (psql *DB) createTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := psql.db.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS yandex (
    		id SERIAL PRIMARY KEY,
   			longurl VARCHAR(255) NOT NULL,
   			shorturl VARCHAR(255) NOT NULL);`)

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

func (psql *DB) Pool() bool {
	return psql.db.Ping(context.Background()) == nil
}
