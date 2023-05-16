package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kisanetik/learn_go_inc1/config"
	"github.com/kisanetik/learn_go_inc1/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	flag.Parse()
	_, port := config.LoadConfig()

	r := chi.NewRouter()

	r.Get("/{id}", handlers.MethodGet)
	r.Post("/", handlers.MethodPost)
	log.Fatal(http.ListenAndServe(port, r))
}
