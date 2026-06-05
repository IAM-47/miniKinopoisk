package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
	"time"
)

type actorRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	BirthDate *string `json:"birth_date,omitempty"`
	Salary    float64 `json:"salary,omitempty"`
}

func parseBirthDate(raw *string) (*time.Time, error) {
	if raw == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", *raw)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func CreateActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req actorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.FirstName == "" || req.LastName == "" {
			http.Error(w, "First name and last name are required", http.StatusBadRequest)
			return
		}

		birthDate, err := parseBirthDate(req.BirthDate)
		if err != nil {
			http.Error(w, "Invalid date format, use yyyy-mm-dd", http.StatusBadRequest)
			return
		}

		actor, err := actorStorage.CreateActor(r.Context(), req.FirstName, req.LastName, birthDate, req.Salary)
		if err != nil {
			log.Printf("CreateActor error: %v", err)
			http.Error(w, "Failed to create actor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(actor)
	}
}

func GetActorsByMovie(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieIDStr := r.PathValue("id")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		actors, err := actorStorage.GetActorsByMovie(r.Context(), movieID)
		if err != nil {
			log.Printf("GetActorsByMovie error: %v", err)
			http.Error(w, "Failed to get actors", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actors)
	}
}

type addActorToMovieRequest struct {
	ActorID int `json:"actor_id"`
}

func AddActorToMovie(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieIDStr := r.PathValue("id")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req addActorToMovieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.ActorID <= 0 {
			http.Error(w, "Invalid actor_id", http.StatusBadRequest)
			return
		}

		if err := actorStorage.AddActorToMovie(r.Context(), movieID, req.ActorID); err != nil {
			log.Printf("AddActorToMovie error: %v", err)
			http.Error(w, "Failed to link actor to movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		var req actorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.FirstName == "" || req.LastName == "" {
			http.Error(w, "First name and last name are required", http.StatusBadRequest)
			return
		}

		birthDate, err := parseBirthDate(req.BirthDate)
		if err != nil {
			http.Error(w, "Invalid date format, use yyyy-mm-dd", http.StatusBadRequest)
			return
		}

		actor, err := actorStorage.UpdateActor(r.Context(), id, req.FirstName, req.LastName, birthDate, req.Salary)
		if err != nil {
			log.Printf("UpdateActor error: %v", err)
			http.Error(w, "Failed to update actor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actor)
	}
}

func DeleteActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		if err := actorStorage.DeleteActor(r.Context(), id); err != nil {
			log.Printf("DeleteActor error: %v", err)
			http.Error(w, "Failed to delete actor", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetActorByID(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		actor, err := actorStorage.GetActorByID(r.Context(), id)
		if err != nil {
			log.Printf("GetActorByID error: %v", err)
			http.Error(w, "Actor not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actor)
	}
}
