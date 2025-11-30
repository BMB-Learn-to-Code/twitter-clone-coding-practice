package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Println(err)
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("error writing JSON response: %v", err))
		return
	}
}
