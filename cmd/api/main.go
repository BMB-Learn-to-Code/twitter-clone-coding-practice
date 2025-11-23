package main

import (
	"log"

	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/env"
	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.Mount()

	log.Panic(app.Run(mux))
}
