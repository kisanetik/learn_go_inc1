package handlers

import (
	"net/http"

	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"github.com/kisanetik/learn_go_inc1/internal/storage/database/postgres"
)

func (a *App) PingHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := postgres.NewPostgresDB(a.Config.DatabaseDSN); err != nil {
		logger.Errorf("error connect to db: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
