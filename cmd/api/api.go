package main

import (
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) Run() error {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    a.config.addr,
		Handler: mux,
	}
	log.Printf("server has started on port: %s", a.config.addr)
	return srv.ListenAndServe()

}
