package main

import (
	"log"
	"net/http"

	"github.com/kisanetik/learn_go_inc1/internal/config"
	"github.com/kisanetik/learn_go_inc1/internal/handlers"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
	store "github.com/kisanetik/learn_go_inc1/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Can't read config: %w", err)
	}

	store, err := store.NewStorage(cfg)
	if err != nil {
		logger.Fatalf("Can't download storage: %w", err)
	}

	app := handlers.NewApp(cfg, store)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, app))
}
