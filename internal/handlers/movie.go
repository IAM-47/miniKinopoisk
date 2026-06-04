package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
)

type movieRequest struct {
	Title       string `json:"title"`
	Producer    string `json:"producer"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year,omitempty"`
}

func CreateMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req movieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Title == "" || req.Producer == "" || req.Director == "" {
			http.Error(w, "Title, Producer and Director are required", http.StatusBadRequest)
			return
		}

		movie, err := movieStorage.CreateMovie(r.Context(), req.Title, req.Producer, req.Director, req.ReleaseYear)
		if err != nil {
			log.Printf("CreateMovie error: %v", err)
			http.Error(w, "Failed to create movie", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movie)
	}
}

func GetMovies(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := 10
		offset := 0
		if l := r.URL.Query().Get("limit"); l != "" {
			if v, err := strconv.Atoi(l); err == nil && v > 0 {
				limit = v
			}
		}
		if o := r.URL.Query().Get("offset"); o != "" {
			if v, err := strconv.Atoi(o); err == nil && v >= 0 {
				offset = v
			}
		}
		movies, err := movieStorage.GetMovies(r.Context(), limit, offset)
		if err != nil {
			log.Printf("GetMovies error: %v", err)
			http.Error(w, "Failed to get movies", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)
	}
}

func GetMovieByID(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		movie, err := movieStorage.GetMovieByID(r.Context(), id)
		if err != nil {
			log.Printf("GetMovieByID error: %v", err)
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie)
	}
}

func UpdateMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req movieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Title == "" || req.Producer == "" || req.Director == "" {
			http.Error(w, "Title, Producer and Director are required", http.StatusBadRequest)
			return
		}

		movie, err := movieStorage.UpdateMovie(r.Context(), id, req.Title, req.Producer, req.Director, req.ReleaseYear)
		if err != nil {
			log.Printf("UpdateMovie error: %v", err)
			http.Error(w, "Failed to update movie", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie)
	}
}

func DeleteMovie(movieStorage *storage.MovieStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		if err := movieStorage.DeleteMovie(r.Context(), id); err != nil {
			log.Printf("DeleteMovie error: %v", err)
			http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
