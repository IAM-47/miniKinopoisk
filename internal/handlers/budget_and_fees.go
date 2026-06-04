package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"net/http"
	"strconv"
)

type budgetRequest struct {
	TotalBudget       float64 `json:"total_budget"`
	FeesInProdCountry float64 `json:"fees_in_prod_country"`
	FeesInOther       float64 `json:"fees_in_other"`
}

func CreateBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req budgetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		budget, err := budgetStorage.CreateBudget(r.Context(), movieID, req.TotalBudget, req.FeesInProdCountry, req.FeesInOther)
		if err != nil {
			log.Printf("CreateBudget error: %v", err)
			http.Error(w, "Failed to create budget", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(budget)
	}
}

func GetBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		budget, err := budgetStorage.GetBudgetByMovie(r.Context(), movieID)
		if err != nil {
			log.Printf("GetBudget error: %v", err)
			http.Error(w, "Budget not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(budget)
	}
}

func UpdateBudgetByMovie(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var req budgetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		budget, err := budgetStorage.UpdateBudgetByMovie(r.Context(), movieID, req.TotalBudget, req.FeesInProdCountry, req.FeesInOther)
		if err != nil {
			log.Printf("UpdateBudget error: %v", err)
			http.Error(w, "Failed to update budget", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(budget)
	}
}

func DeleteBudget(budgetStorage *storage.BudgetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid budget ID", http.StatusBadRequest)
			return
		}

		if err := budgetStorage.DeleteBudget(r.Context(), id); err != nil {
			log.Printf("DeleteBudget error: %v", err)
			http.Error(w, "Failed to delete budget", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
