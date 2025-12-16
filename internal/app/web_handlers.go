package app

import (
	"log"
	"net/http"
	"strconv"

	"miniKinopoisk/internal/auth"
	"miniKinopoisk/internal/models"
	"miniKinopoisk/internal/storage"
	"miniKinopoisk/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

func (app *App) getUserFromCookie(r *http.Request) *models.User {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil
	}

	claims := &auth.Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mini-kinopoisk-secret-key"), nil
	})
	if err != nil || !token.Valid {
		return nil
	}

	return &models.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}
}

func (app *App) render(w http.ResponseWriter, page string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := app.templates.ExecuteTemplate(w, page+".html", data); err != nil {
		log.Printf("Ошибка рендеринга %s.html: %v", page, err)
		http.Error(w, "500 — ошибка шаблона", http.StatusInternalServerError)
	}
}

func (app *App) handleHome(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	movieStorage := storage.NewMovieStorage(app.db)
	movies, _ := movieStorage.GetMovies(r.Context())
	app.render(w, "index", struct {
		User   *models.User
		Movies []*models.Movie
	}{
		User:   user,
		Movies: movies,
	})
}

func (app *App) handleMovie(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	movieIDStr := r.PathValue("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movieStorage := storage.NewMovieStorage(app.db)
	movie, err := movieStorage.GetMovieByID(r.Context(), movieID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	actorStorage := storage.NewActorStorage(app.db)
	actors, _ := actorStorage.GetActorsByMovie(r.Context(), movieID)

	budgetStorage := storage.NewBudgetStorage(app.db)
	budget, _ := budgetStorage.GetBudgetByMovie(r.Context(), movieID)

	app.render(w, "movie", struct {
		User   *models.User
		Movie  *models.Movie
		Actors []*models.Actor
		Budget *models.BudgetAndFees
	}{
		User:   user,
		Movie:  movie,
		Actors: actors,
		Budget: budget,
	})
}

func (app *App) handleLogin(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	app.render(w, "login", struct {
		User   *models.User
		Title  string
		Action string
	}{
		User:   user,
		Title:  "Вход",
		Action: "/login",
	})
}

func (app *App) handleLoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	userStorage := storage.NewUserStorage(app.db)
	user, err := userStorage.GetUserByEmail(r.Context(), email)
	if err != nil || !utils.CheckPasswordHash(password, user.PasswordHash) {
		http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600 * 24,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *App) handleRegister(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	app.render(w, "login", struct {
		User   *models.User
		Title  string
		Action string
	}{
		User:   user,
		Title:  "Регистрация",
		Action: "/register",
	})
}

func (app *App) handleRegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Ошибка хеширования", http.StatusInternalServerError)
		return
	}

	userStorage := storage.NewUserStorage(app.db)
	_, err = userStorage.CreateUser(r.Context(), email, hash)
	if err != nil {
		http.Error(w, "Пользователь уже существует", http.StatusConflict)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (app *App) handleCreateMovieForm(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	if user == nil || user.Role != "admin" {
		http.Error(w, "Доступ запрещён", http.StatusForbidden)
		return
	}
	app.render(w, "create_movie", nil)
}

func (app *App) handleCreateMoviePost(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromCookie(r)
	if user == nil || user.Role != "admin" {
		http.Error(w, "Доступ запрещён", http.StatusForbidden)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	producer := r.FormValue("producer")
	director := r.FormValue("director")
	releaseYear, _ := strconv.Atoi(r.FormValue("release_year"))

	movieStorage := storage.NewMovieStorage(app.db)
	_, err := movieStorage.CreateMovie(r.Context(), title, producer, director, releaseYear)
	if err != nil {
		http.Error(w, "Ошибка создания фильма", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
