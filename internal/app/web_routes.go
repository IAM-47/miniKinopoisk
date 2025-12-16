package app

import (
	"miniKinopoisk/internal/middleware"
	"net/http"
)

func (app *App) registerWebRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", app.handleHome)
	mux.HandleFunc("GET /movies/{id}", app.handleMovie)

	mux.HandleFunc("GET /login", app.handleLogin)
	mux.HandleFunc("POST /login", app.handleLoginPost)

	mux.HandleFunc("GET /register", app.handleRegister)
	mux.HandleFunc("POST /register", app.handleRegisterPost)

	mux.HandleFunc("GET /movies/create", app.handleCreateMovieForm)
	mux.HandleFunc("POST /movies/create", middleware.AuthMiddleware(middleware.AdminOnly(app.handleCreateMoviePost)))
}
