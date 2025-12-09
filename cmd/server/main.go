package main

import (
	"log"
	"net/http"

	"miniKinopoisk/internal/handlers"
	"miniKinopoisk/pkg/config"
)

func main() {
	// Загружаем конфиг
	cfg := config.Load()

	// Создаём маршрутизатор (пока простой)
	router := handlers.NewRouter(cfg)

	// Запускаем сервер
	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
