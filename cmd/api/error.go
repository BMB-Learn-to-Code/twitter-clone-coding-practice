package main

import (
	"log"
	"net/http"
)

func (a *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%d interna server error: %s, path: %s, error: %s", http.StatusInternalServerError, r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem.")
}
