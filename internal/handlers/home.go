package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
)

var version = "1.0.0"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"app":     "miniKinopoisk API",
		"version": version,
		"status":  "ok",
		"go":      runtime.Version(),
	})
}
