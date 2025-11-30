package main

import "net/http"

func writeJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	return nil
}
