package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"miniKinopoisk/internal/models"
)

var errBoom = errors.New("boom")

// serve поднимает хендлер в настоящем ServeMux с паттерном маршрута,
// чтобы r.PathValue("id") работал так же, как в проде.
func serve(pattern string, h http.HandlerFunc, method, target, body string) *httptest.ResponseRecorder {
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, h)

	var req *http.Request
	if body == "" {
		req = httptest.NewRequest(method, target, nil)
	} else {
		req = httptest.NewRequest(method, target, strings.NewReader(body))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	return rec
}

// ---- fakeUsers ----

type fakeUsers struct {
	createFn func(ctx context.Context, email, hash string) (*models.User, error)
	getFn    func(ctx context.Context, email string) (*models.User, error)
	deleteFn func(ctx context.Context, email string) error
}

func (f *fakeUsers) CreateUser(ctx context.Context, email, hash string) (*models.User, error) {
	return f.createFn(ctx, email, hash)
}
func (f *fakeUsers) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return f.getFn(ctx, email)
}
func (f *fakeUsers) DeleteUserByEmail(ctx context.Context, email string) error {
	return f.deleteFn(ctx, email)
}

// ---- fakeMovies ----

type fakeMovies struct {
	createFn  func(title, producer, director string, year int) (*models.Movie, error)
	listFn    func(limit, offset int) ([]*models.Movie, error)
	getFn     func(id int) (*models.Movie, error)
	updateFn  func(id int, title, producer, director string, year int) (*models.Movie, error)
	deleteFn  func(id int) error
	byActorFn func(actorID int) ([]*models.Movie, error)
}

func (f *fakeMovies) CreateMovie(_ context.Context, t, p, d string, y int) (*models.Movie, error) {
	return f.createFn(t, p, d, y)
}
func (f *fakeMovies) GetMovies(_ context.Context, limit, offset int) ([]*models.Movie, error) {
	return f.listFn(limit, offset)
}
func (f *fakeMovies) GetMovieByID(_ context.Context, id int) (*models.Movie, error) {
	return f.getFn(id)
}
func (f *fakeMovies) UpdateMovie(_ context.Context, id int, t, p, d string, y int) (*models.Movie, error) {
	return f.updateFn(id, t, p, d, y)
}
func (f *fakeMovies) DeleteMovie(_ context.Context, id int) error { return f.deleteFn(id) }
func (f *fakeMovies) GetMoviesByActor(_ context.Context, actorID int) ([]*models.Movie, error) {
	return f.byActorFn(actorID)
}

// ---- fakeActors ----

type fakeActors struct {
	createFn  func(first, last string, birth *time.Time, salary float64) (*models.Actor, error)
	byMovieFn func(movieID int) ([]*models.Actor, error)
	linkFn    func(movieID, actorID int) error
	updateFn  func(id int, first, last string, birth *time.Time, salary float64) (*models.Actor, error)
	deleteFn  func(id int) error
	getFn     func(id int) (*models.Actor, error)
}

func (f *fakeActors) CreateActor(_ context.Context, first, last string, birth *time.Time, salary float64) (*models.Actor, error) {
	return f.createFn(first, last, birth, salary)
}
func (f *fakeActors) GetActorsByMovie(_ context.Context, movieID int) ([]*models.Actor, error) {
	return f.byMovieFn(movieID)
}
func (f *fakeActors) AddActorToMovie(_ context.Context, movieID, actorID int) error {
	return f.linkFn(movieID, actorID)
}
func (f *fakeActors) UpdateActor(_ context.Context, id int, first, last string, birth *time.Time, salary float64) (*models.Actor, error) {
	return f.updateFn(id, first, last, birth, salary)
}
func (f *fakeActors) DeleteActor(_ context.Context, id int) error { return f.deleteFn(id) }
func (f *fakeActors) GetActorByID(_ context.Context, id int) (*models.Actor, error) {
	return f.getFn(id)
}

// ---- fakeBudgets ----

type fakeBudgets struct {
	createFn func(movieID int, total, prod, other float64) (*models.BudgetAndFees, error)
	getFn    func(movieID int) (*models.BudgetAndFees, error)
	updateFn func(movieID int, total, prod, other float64) (*models.BudgetAndFees, error)
	deleteFn func(id int) error
}

func (f *fakeBudgets) CreateBudget(_ context.Context, movieID int, total, prod, other float64) (*models.BudgetAndFees, error) {
	return f.createFn(movieID, total, prod, other)
}
func (f *fakeBudgets) GetBudgetByMovie(_ context.Context, movieID int) (*models.BudgetAndFees, error) {
	return f.getFn(movieID)
}
func (f *fakeBudgets) UpdateBudgetByMovie(_ context.Context, movieID int, total, prod, other float64) (*models.BudgetAndFees, error) {
	return f.updateFn(movieID, total, prod, other)
}
func (f *fakeBudgets) DeleteBudget(_ context.Context, id int) error { return f.deleteFn(id) }
