package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kisanetik/learn_go_inc1/config"
	"github.com/kisanetik/learn_go_inc1/internal/handlers"
	logger "github.com/kisanetik/learn_go_inc1/internal/logging"

	"github.com/go-chi/chi/v5"
)

func main() {
	flag.Parse()
	_, port := config.LoadConfig()

	r := chi.NewRouter()

	r.Get("/{id}", logger.ResponseLogger(handlers.MethodGet))
	r.Post("/", logger.RequestLogger(handlers.MethodPost))
	r.Post("/api/shorten", logger.RequestLogger(handlers.JsonPost))
	log.Fatal(http.ListenAndServe(port, r))
}
