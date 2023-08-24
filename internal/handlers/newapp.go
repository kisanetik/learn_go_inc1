package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kisanetik/learn_go_inc1/internal/config"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"github.com/kisanetik/learn_go_inc1/internal/middleware/gzip"
	"github.com/kisanetik/learn_go_inc1/internal/storage"
)

type App struct {
	*chi.Mux
	Config  config.Config
	Storage storage.Storage
}

func NewApp(cfg config.Config, store storage.Storage) *App {
	app := App{
		chi.NewRouter(),
		cfg,
		store,
	}

	app.Use(gzip.Request)
	app.Use(gzip.Response)

	app.Route("/", func(r chi.Router) {
		r.Post("/", logger.RequestLogger(app.CompressHandler))
		r.Get("/{id}", logger.ResponseLogger(app.GetURLHandler))
		r.Get("/ping", logger.ResponseLogger(app.PingHandler))

		r.Route("/api", func(r chi.Router) {
			r.Post("/shorten", app.JSONHandler)
		})
	})

	return &app
}
