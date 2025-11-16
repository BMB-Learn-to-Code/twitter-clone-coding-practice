package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.Mount()

	log.Panic(app.Run(mux))
}
