package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
	"time"
)

type createActorRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	BirthDate *string `json:"birth_date,omitempty"`
	Salary    float64 `json:"salary,omitempty"`
}

func CreateActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createActorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.FirstName == "" || req.LastName == "" {
			http.Error(w, "First Name and Last Name are required", http.StatusBadRequest)
			return
		}

		var birthDate *time.Time
		if req.BirthDate != nil {
			t, err := time.Parse("2006-01-02", *req.BirthDate)
			if err != nil {
				http.Error(w, "invalid date format, needed yyyy-mm-dd", http.StatusBadRequest)
				return
			}
			birthDate = &t
		}

		actor, err := actorStorage.CreateActor(r.Context(), req.FirstName, req.LastName, birthDate, req.Salary)
		if err != nil {
			http.Error(w, "Something wend wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(actor)
	}
}

func GetActors(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actors, err := actorStorage.GetActors(r.Context())
		if err != nil {
			http.Error(w, "Something wend wrong", http.StatusInternalServerError)
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
		if movieIDStr == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req addActorToMovieRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.ActorID <= 0 {
			http.Error(w, "Invalid actorID", http.StatusBadRequest)
			return
		}
		if err := actorStorage.AddActorToMovie(r.Context(), movieID, req.ActorID); err != nil {
			http.Error(w, "Something wend wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type updateActorRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	BirthDate *string `json:"birth_date,omitempty"`
	Salary    float64 `json:"salary,omitempty"`
}

func UpdateActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		if id_str == "" {
			http.Error(w, "Actor ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		var req updateActorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if req.FirstName == "" || req.LastName == "" {
			http.Error(w, "FirstName, Lastname are required", http.StatusBadRequest)
			return
		}
		var birthDate *time.Time
		if req.BirthDate != nil {
			t, err := time.Parse("2006-01-02", *req.BirthDate)
			if err != nil {
				http.Error(w, "Invalid date format, use yyyy-mm-dd", http.StatusBadRequest)
				return
			}
			birthDate = &t
		}

		actor, err := actorStorage.UpdateActor(r.Context(), id, req.FirstName, req.LastName, birthDate, req.Salary)
		if err != nil {
			log.Printf("Something wend wrong, %v", err)
			http.Error(w, "Error updating actor: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actor)
	}
}

func DeleteActor(actorStorage *storage.ActorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		if id_str == "" {
			http.Error(w, "Actor ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		if err := actorStorage.DeleteActor(r.Context(), id); err != nil {
			http.Error(w, "Error deleting actor", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
