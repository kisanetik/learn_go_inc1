package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	urlmaker "github.com/kisanetik/learn_go_inc1/internal/app"
	"github.com/kisanetik/learn_go_inc1/internal/storage"
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
	data := storage.GetData()
	hash := strings.TrimPrefix(req.RequestURI, "/")
	if val, ok := data[hash]; ok {
		res.Header().Add("Location", string(val.OriginalURL))
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusNotFound)
	}
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
	jsonres.ShortURL = urlmaker.CompressURL(string(url.UserURL))
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
