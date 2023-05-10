package main

import (
	"log"
	"net/http"

	"github.com/kisanetik/learn_go_inc1/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/{id}", handlers.MethodGet)
	r.Post("/", handlers.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
