package app

import (
	"miniKinopoisk/internal/handlers"
	"miniKinopoisk/internal/middleware"
	"miniKinopoisk/internal/storage"
	"net/http"
)

func (app *App) registerAPIRoutes(mux *http.ServeMux) {
	moviesStorage := storage.NewMovieStorage(app.db)
	actorsStorage := storage.NewActorStorage(app.db)
	budgetStorage := storage.NewBudgetStorage(app.db)

	// Фильмы
	mux.HandleFunc("GET /movies", handlers.GetMovies(moviesStorage))
	mux.HandleFunc("POST /movies", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateMovie(moviesStorage))))
	mux.HandleFunc("PUT /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateMovie(moviesStorage))))
	mux.HandleFunc("DELETE /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteMovie(moviesStorage))))

	// Актёры
	mux.HandleFunc("GET /movies/{id}/actors", handlers.GetActorsByMovie(actorsStorage))
	mux.HandleFunc("POST /actors", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateActor(actorsStorage))))
	mux.HandleFunc("PUT /actors/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateActor(actorsStorage))))
	mux.HandleFunc("DELETE /actors/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteActor(actorsStorage))))

	// Привязка актёра к фильму
	mux.HandleFunc("POST /movies/{id}/actors", middleware.AuthMiddleware(middleware.AdminOnly(handlers.AddActorToMovie(actorsStorage))))

	// Бюджет и сборы
	mux.HandleFunc("GET /movies/{id}/budget", handlers.GetBudget(budgetStorage))
	mux.HandleFunc("POST /budget", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateBudget(budgetStorage))))
	mux.HandleFunc("PUT /movies/{id}/budget", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateBudgetByMovie(budgetStorage))))
	mux.HandleFunc("DELETE /budget/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteBudget(budgetStorage))))
}
