package handlers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	linker "github.com/kisanetik/learn_go_inc1/internal/app"
)

func MethodPost(res http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	if len(body) == 0 {
		res.WriteHeader(http.StatusBadRequest)
	}
	location := linker.CompressURL(string(body))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(location))
}

func MethodGet(res http.ResponseWriter, req *http.Request) {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	filename := filepath.Dir(tFile.Name()) + req.RequestURI
	if _, err := os.Stat(filename); err == nil {
		data, fErr := os.ReadFile(filename)
		if fErr != nil {
			panic(fErr)
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

func LinkerHandler(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	if request.Method == http.MethodPost {
		MethodPost(res, request)
	} else if request.Method == http.MethodGet {
		MethodGet(res, request)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
