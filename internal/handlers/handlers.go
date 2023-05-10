package handlers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func methodPost(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	tFile, err := os.CreateTemp("", "")
	os.WriteFile(tFile.Name(), body, 0644)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(tFile.Name(), body, 0644)
	if err != nil {
		panic(err)
	}
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("http://localhost:8080/" + filepath.Base(tFile.Name())))
}

func methodGet(res http.ResponseWriter, req *http.Request) {
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

func LinkerHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")

	if request.Method == http.MethodPost {
		methodPost(writer, request)
	} else if request.Method == http.MethodGet {
		methodGet(writer, request)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}
