package main

import (
	"context"
	"encoding/json"
	"log"
	"miniKinopoisk/internal/middleware"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"miniKinopoisk/internal/handlers"
	"miniKinopoisk/internal/storage"
	"miniKinopoisk/pkg/config"
)

// Простой хендлер для корня
func homeHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"message": "Mini Kinopoisk API",
		"status":  "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	cfg := config.Load()

	db, err := pgxpool.New(context.Background(), cfg.Database.DSN)
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer db.Close()

	userStorage := storage.NewUserStorage(db)
	movieStorage := storage.NewMovieStorage(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler) // ← добавили!
	// Пользователи
	mux.HandleFunc("POST /register", handlers.Register(userStorage))
	mux.HandleFunc("POST /login", handlers.Login(userStorage))

	// Фильмы
	mux.HandleFunc("POST /movies", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateMovie(movieStorage))))
	mux.HandleFunc("PUT /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateMovie(movieStorage))))
	mux.HandleFunc("DELETE /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteMovie(movieStorage))))
	/// Читать могут все
	mux.HandleFunc("GET /movies", handlers.GetMovies(movieStorage))
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
