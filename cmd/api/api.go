package main

import (
	"log"
	"net/http"
	"time"
	"twitter-clone-coding-practice/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns uint32
	maxIdleConns uint32
	maxIdleTime  string
}

func (a *application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", a.createPostHandler)
			r.Get("/{id}", a.getPostHandler)
		})
	})

	return r
}

func (a *application) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started on port: %s", a.config.addr)
	return srv.ListenAndServe()

}
