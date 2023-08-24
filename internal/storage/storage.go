package storage

import (
	"fmt"
	"github.com/kisanetik/learn_go_inc1/internal/config"
	"github.com/kisanetik/learn_go_inc1/internal/storage/fs"
	"github.com/kisanetik/learn_go_inc1/internal/storage/mem"
	"os"
)

type Storage interface {
	Save(string) (string, error)
	Get(string) string
	Close() error
}

func NewStorage(cfg config.Config) (Storage, error) {
	var s Storage
	var err error

	if path := cfg.FileStoragePath; path != "" {
		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			return nil, fmt.Errorf("cannot open file: %w", err)
		}

		if s, err = fs.NewFs(file); err != nil {
			return nil, fmt.Errorf("error NewFs file: %w", err)
		}
	} else {
		if s, err = mem.NewMem(); err != nil {
			return nil, fmt.Errorf("error NewMem: %w", err)
		}
	}

	return s, nil
}
