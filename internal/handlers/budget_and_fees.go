package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
)

type createBudgetRequest struct {
	IDMovie           int     `json:"id_movie"`
	TotalBudget       float64 `json:"total_budget,omitempty"`
	FeesInProdCountry float64 `json:"fees_in_prod_country,omitempty"`
	FeesInOther       float64 `json:"fees_in_other,omitempty"`
}

func CreateBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBudgetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.IDMovie == 0 {
			http.Error(w, "The movie id is required", http.StatusBadRequest)
			return
		}

		budget, err := budgetStorage.CreateBudget(r.Context(), req.IDMovie, req.TotalBudget, req.FeesInProdCountry, req.FeesInOther)
		if err != nil {
			http.Error(w, "Something wend wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(budget)
	}
}

func GetBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieIDStr := r.PathValue("id")
		if movieIDStr == "" {
			http.Error(w, "The movie id is required", http.StatusBadRequest)
			return
		}

		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}
		budget, err := budgetStorage.GetBudgetByMovie(r.Context(), movieID)
		if err != nil {
			http.Error(w, "Something wend wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(budget)
	}
}

type updateBudgetRequest struct {
	IDMovie           int     `json:"id_movie"`
	TotalBudget       float64 `json:"total_budget,omitempty"`
	FeesInProdCountry float64 `json:"fees_in_prod_country,omitempty"`
	FeesInOther       float64 `json:"fees_in_other,omitempty"`
}

func UpdateBudgetByMovie(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
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

		var req updateBudgetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if req.IDMovie == 0 {
			http.Error(w, "The Movie ID is required", http.StatusBadRequest)
			return
		}

		budget, err := budgetStorage.UpdateBudgetByMovie(r.Context(), movieID, req.TotalBudget, req.FeesInProdCountry, req.FeesInOther)
		if err != nil {
			log.Printf("Something wend wrong, %v", err)
			http.Error(w, "Error updating budget: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(budget)
	}
}

func DeleteBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			http.Error(w, "Budget ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid budget ID", http.StatusBadRequest)
			return
		}

		if err := budgetStorage.DeleteBudget(r.Context(), id); err != nil {
			http.Error(w, "Error deleting budget", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
