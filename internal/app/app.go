package app

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	db        *pgxpool.Pool
	templates *template.Template
}

func NewApp(db *pgxpool.Pool) *App {
	dir, _ := os.Getwd()
	log.Println("dir:", dir)
	log.Println("Templates in:", dir+"/web/templates/")

	tmpl := template.Must(template.ParseFiles(
		"./web/templates/base.html",
		"./web/templates/index.html",
		"./web/templates/movie.html",
		"./web/templates/login.html",
		"./web/templates/create_movie.html",
	))
	return &App{
		db:        db,
		templates: tmpl,
	}
}

func (app *App) RegisterRoutes(mux *http.ServeMux) {
	app.registerWebRoutes(mux)
	app.registerAPIRoutes(mux)
}
