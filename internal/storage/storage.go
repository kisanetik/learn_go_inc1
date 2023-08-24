package storage

import (
	"fmt"

	"github.com/kisanetik/learn_go_inc1/internal/config"
	"github.com/kisanetik/learn_go_inc1/internal/storage/database/postgres"
	"github.com/kisanetik/learn_go_inc1/internal/storage/fs"
	"github.com/kisanetik/learn_go_inc1/internal/storage/mem"
)

type Storage interface {
	Save(string) (string, error)
	Get(string) string
	Close() error
}

func NewStorage(cfg config.Config) (Storage, error) {
	var s Storage
	var err error

	if cfg.DatabaseDSN != "" {
		if s, err = postgres.NewPostgresDB(cfg.DatabaseDSN); err != nil {
			return nil, fmt.Errorf("Can't database storage: %w", err)
		}
	} else if cfg.FileStoragePath != "" {
		if s, err = fs.NewFsFromFile(cfg.FileStoragePath); err != nil {
			return nil, fmt.Errorf("Error NewFs file: %w", err)
		}
	} else {
		if s, err = mem.NewMem(); err != nil {
			return nil, fmt.Errorf("Error NewMem: %w", err)
		}
	}

	return s, nil
}
