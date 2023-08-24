package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/kisanetik/learn_go_inc1/internal/logger"
)

type URLRequest struct {
	ServerURL string `json:"url"`
}

type ResultResponse struct {
	BaseShortURL string `json:"result"`
}

func (a *App) JSONHandler(w http.ResponseWriter, r *http.Request) {
	var req URLRequest
	var resp ResultResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("json decode error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := a.Storage.Save(req.ServerURL)
	if err != nil {
		logger.Errorf("Storage save error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.BaseShortURL, err = url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		logger.Errorf("Join err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respContent, err := json.Marshal(resp)
	if err != nil {
		logger.Errorf("JSON marshal error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(respContent); err != nil {
		logger.Errorf("Failed to send URL on json handler: %s", err)
	}
}
