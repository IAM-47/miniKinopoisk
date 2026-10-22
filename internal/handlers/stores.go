package handlers

import (
	"context"
	"time"

	"miniKinopoisk/internal/models"
)

// Интерфейсы объявлены на стороне потребителя (хендлеров): так обработчики
// не зависят от pgx и их можно тестировать на подменных реализациях без БД.
// Реальные *storage.XxxStorage удовлетворяют им без изменений.

type UserStore interface {
	CreateUser(ctx context.Context, email, hash string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
}

type MovieStore interface {
	CreateMovie(ctx context.Context, title, producer, director string, releaseYear int) (*models.Movie, error)
	GetMovies(ctx context.Context, limit, offset int) ([]*models.Movie, error)
	GetMovieByID(ctx context.Context, id int) (*models.Movie, error)
	UpdateMovie(ctx context.Context, id int, title, producer, director string, releaseYear int) (*models.Movie, error)
	DeleteMovie(ctx context.Context, id int) error
	GetMoviesByActor(ctx context.Context, actorID int) ([]*models.Movie, error)
}

type ActorStore interface {
	CreateActor(ctx context.Context, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error)
	GetActorsByMovie(ctx context.Context, movieID int) ([]*models.Actor, error)
	AddActorToMovie(ctx context.Context, movieID, actorID int) error
	UpdateActor(ctx context.Context, id int, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error)
	DeleteActor(ctx context.Context, id int) error
	GetActorByID(ctx context.Context, id int) (*models.Actor, error)
}

type BudgetStore interface {
	CreateBudget(ctx context.Context, movieID int, totalBudget, feesInProdCountry, feesInOther float64) (*models.BudgetAndFees, error)
	GetBudgetByMovie(ctx context.Context, movieID int) (*models.BudgetAndFees, error)
	UpdateBudgetByMovie(ctx context.Context, movieID int, totalBudget, feesInProdCountry, feeInOther float64) (*models.BudgetAndFees, error)
	DeleteBudget(ctx context.Context, id int) error
}
