package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kisanetik/learn_go_inc1/config"
	"github.com/kisanetik/learn_go_inc1/internal/handlers"
	gzip "github.com/kisanetik/learn_go_inc1/internal/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	flag.Parse()
	_, port := config.LoadConfig()

	r := chi.NewRouter()

	//MIDDLEWARE LIST
	r.Use(middleware.Logger)
	r.Use(gzip.Request)
	r.Use(gzip.Response)

	//ROUTES LIST
	r.Get("/{id}", handlers.MethodGet)
	r.Post("/", handlers.MethodPost)
	r.Post("/api/shorten", handlers.JSONsPost)

	log.Fatal(http.ListenAndServe(port, r))
}
