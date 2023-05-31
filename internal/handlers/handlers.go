package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	urlmaker "github.com/kisanetik/learn_go_inc1/internal/app"
)

func MethodPost(res http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	if len(body) == 0 {
		res.WriteHeader(http.StatusBadRequest)
	}
	location := urlmaker.CompressURL(string(body))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(location))
}

func MethodGet(res http.ResponseWriter, req *http.Request) {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	filename := filepath.Dir(tFile.Name()) + req.RequestURI
	if _, err := os.Stat(filename); err == nil {
		data, fErr := os.ReadFile(filename)
		if fErr != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Header().Add("Location", string(data))
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else if errors.Is(err, os.ErrNotExist) {
		res.WriteHeader(http.StatusNotFound)
	} else {
		panic(err)
	}
	defer os.Remove(tFile.Name())
}

type URL struct {
	UserURL string `json:"url"`
}

type Result struct {
	ShortURL string `json:"result"`
}

func JSONsPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var url URL
	var jsonres Result
	if err := json.NewDecoder(req.Body).Decode(&url); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var err error
	jsonres.ShortURL = urlmaker.CompressURL(string(url.UserURL))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(map[string]string{"result": jsonres.ShortURL})
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		log.Fatal("Failed to send URL on json handler")
	}
}
