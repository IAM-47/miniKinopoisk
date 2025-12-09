package handlers

import (
	"encoding/json"
	"net/http"

	"miniKinopoisk/pkg/config"
)

type Router struct {
	cfg *config.Config
}

func NewRouter(cfg *config.Config) *http.ServeMux {
	router := http.NewServeMux()
	handler := &Router{cfg: cfg}

	router.HandleFunc("GET /", handler.Home)

	return router
}

func (r *Router) Home(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{
		"message": "Mini Kinopoisk is ready!",
		"port":    r.cfg.Server.Port,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
