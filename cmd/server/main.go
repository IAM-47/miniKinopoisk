package main

import (
	"context"
	"encoding/json"
	"log"
	"miniKinopoisk/internal/app"
	"miniKinopoisk/pkg/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Mini Kinopoisk API",
		"status":  "OK",
	})
}
func main() {
	cfg := config.Load()

	db, err := pgxpool.New(context.Background(), cfg.Database.DSN)
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	application := app.NewApp(db)
	application.RegisterRoutes(mux)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
