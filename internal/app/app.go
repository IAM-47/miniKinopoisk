package app

import (
	"net/http"

	"miniKinopoisk/internal/handlers"
	"miniKinopoisk/internal/middleware"
	"miniKinopoisk/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	db *pgxpool.Pool
}

func NewApp(db *pgxpool.Pool) *App {
	return &App{db: db}
}

func (app *App) RegisterRoutes(mux *http.ServeMux) {
	userStorage := storage.NewUserStorage(app.db)
	moviesStorage := storage.NewMovieStorage(app.db)
	actorsStorage := storage.NewActorStorage(app.db)
	budgetStorage := storage.NewBudgetStorage(app.db)

	mux.HandleFunc("GET /", handlers.HomeHandler)

	// Пользователи
	mux.HandleFunc("POST /register", handlers.Register(userStorage))
	mux.HandleFunc("POST /login", handlers.Login(userStorage))
	mux.HandleFunc("DELETE /users/{email}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteUserByEmail(userStorage))))

	// Фильмы
	mux.HandleFunc("GET /movies", handlers.GetMovies(moviesStorage))
	mux.HandleFunc("GET /movies/{id}", handlers.GetMovieByID(moviesStorage))
	mux.HandleFunc("POST /movies", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateMovie(moviesStorage))))
	mux.HandleFunc("PUT /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateMovie(moviesStorage))))
	mux.HandleFunc("DELETE /movies/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteMovie(moviesStorage))))
	mux.HandleFunc("GET /actors/{id}/movies", handlers.GetMoviesByActor(moviesStorage))

	// Актёры
	mux.HandleFunc("GET /movies/{id}/actors", handlers.GetActorsByMovie(actorsStorage))
	mux.HandleFunc("GET /actors/{id}", handlers.GetActorByID(actorsStorage))
	mux.HandleFunc("POST /actors", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateActor(actorsStorage))))
	mux.HandleFunc("PUT /actors/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateActor(actorsStorage))))
	mux.HandleFunc("DELETE /actors/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteActor(actorsStorage))))
	mux.HandleFunc("POST /movies/{id}/actors", middleware.AuthMiddleware(middleware.AdminOnly(handlers.AddActorToMovie(actorsStorage))))

	// Бюджет и сборы
	mux.HandleFunc("GET /movies/{id}/budget", handlers.GetBudget(budgetStorage))
	mux.HandleFunc("POST /movies/{id}/budget", middleware.AuthMiddleware(middleware.AdminOnly(handlers.CreateBudget(budgetStorage))))
	mux.HandleFunc("PUT /movies/{id}/budget", middleware.AuthMiddleware(middleware.AdminOnly(handlers.UpdateBudgetByMovie(budgetStorage))))
	mux.HandleFunc("DELETE /budget/{id}", middleware.AuthMiddleware(middleware.AdminOnly(handlers.DeleteBudget(budgetStorage))))
}
