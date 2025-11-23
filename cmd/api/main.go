package main

import (
	"log"

	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/env"
	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost:5432/twitter_clone?sslmode=disable"),
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxConnLifetime: env.GetString("DB_MAX_CONN_LIFETIME", "15min"),
		},
	}

	store := store.NewStorage(nil

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.Mount()

	log.Panic(app.Run(mux))
}
