package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "id")
	if url == "" {
		_ = errors.New("url param bad with id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := a.Storage.Get(url)
	if m == "" {
		logger.Errorf("get url is bad: %s", url)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
