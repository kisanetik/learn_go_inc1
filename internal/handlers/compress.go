package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/kisanetik/learn_go_inc1/internal/logger"
)

func (a *App) CompressHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		logger.Errorf("body is nil or empty: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if a.Config.DatabaseDSN != "" {
		short, err := a.Storage.CheckIsURLExists(string(body))
		if err != nil {
			logger.Errorf("error is CheckIsURLExists: %s", err)
		}

		if short != "" {
			res, err := url.JoinPath(a.Config.BaseShortURL, short)
			if err != nil {
				logger.Errorf("error is JoinPath: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusConflict)
			_, err = w.Write([]byte(res))
			if err != nil {
				logger.Errorf("Failed to send URL: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			return
		}
	}

	short, err := a.Storage.Save(string(body), "")
	if err != nil {
		logger.Errorf("storage save is error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	long, err := url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		logger.Errorf("join path have err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(long))
	if err != nil {
		logger.Errorf("Failed to send URL: %s", err)
	}
}
