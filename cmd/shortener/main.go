package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		tFile, err := os.CreateTemp("", "")
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(tFile.Name(), body, 0644)
		if err != nil {
			panic(err)
		}
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte("http://localhost:8080/" + filepath.Base(tFile.Name())))
	} else {
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
}

func main() {

	http.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}
