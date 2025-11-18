package main

import (
	"log"

	"github.com/BMB-Learn-to-Code/twitter-clone-coding-practice/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8081"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.Mount()

	log.Panic(app.Run(mux))
}
