package handlers

import (
	"encoding/json"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
)

type createMovieRequest struct {
	Title       string `json:"title"`
	Producer    string `json:"producer"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year,omitempty"`
}

func CreateMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMovieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if req.Title == "" || req.Producer == "" || req.Director == "" {
			http.Error(w, "Title, Producer and Director are required", http.StatusBadRequest)
			return
		}

		movie, err := movieStorage.CreateMovie(r.Context(), req.Title, req.Producer, req.Director, req.ReleaseYear)
		if err != nil {
			http.Error(w, "Error creating movie", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movie)
	}
}

func GetMovies(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := movieStorage.GetMovies(r.Context())
		if err != nil {
			http.Error(w, "Error getting movies", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)
	}
}

type updateMovieRequest struct {
	Title       string `json:"title"`
	Producer    string `json:"producer"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year,omitempty"`
}

func UpdateMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		if id_str == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req updateMovieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if req.Title == "" || req.Producer == "" || req.Director == "" {
			http.Error(w, "Title, Producer and Director are required", http.StatusBadRequest)
			return
		}
		movie, err := movieStorage.UpdateMovie(r.Context(), id, req.Title, req.Producer, req.Director, req.ReleaseYear)
		if err != nil {
			http.Error(w, "Error updating movie", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie)
	}
}

func DeleteMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		if id_str == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		if err := movieStorage.DeleteMovie(r.Context(), id); err != nil {
			http.Error(w, "Error deleting movie", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
