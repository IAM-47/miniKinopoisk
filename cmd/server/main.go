package main

import (
	"context"
	"log"
	"miniKinopoisk/internal/app"
	"miniKinopoisk/pkg/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

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
