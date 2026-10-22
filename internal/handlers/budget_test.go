package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"miniKinopoisk/internal/models"
)

func TestCreateBudget(t *testing.T) {
	var gotMovie int
	var gotTotal float64
	store := &fakeBudgets{createFn: func(movieID int, total, prod, other float64) (*models.BudgetAndFees, error) {
		gotMovie, gotTotal = movieID, total
		return &models.BudgetAndFees{ID: 1, IDMovie: movieID, TotalBudget: total, FeesInProdCountry: prod, FeesInOther: other}, nil
	}}

	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{"success", "/movies/3/budget", `{"total_budget":165000000,"fees_in_prod_country":108000000,"fees_in_other":295000000}`, http.StatusCreated},
		{"invalid movie id", "/movies/x/budget", `{}`, http.StatusBadRequest},
		{"invalid body", "/movies/3/budget", `{`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("POST /movies/{id}/budget", CreateBudget(store), http.MethodPost, tt.target, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
	if gotMovie != 3 || gotTotal != 165000000 {
		t.Errorf("storage got movie=%d total=%v", gotMovie, gotTotal)
	}

	failing := &fakeBudgets{createFn: func(int, float64, float64, float64) (*models.BudgetAndFees, error) { return nil, errBoom }}
	rec := serve("POST /movies/{id}/budget", CreateBudget(failing), http.MethodPost, "/movies/3/budget", `{}`)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("storage failure status = %d, want 500", rec.Code)
	}
}

func TestGetBudget(t *testing.T) {
	store := &fakeBudgets{getFn: func(movieID int) (*models.BudgetAndFees, error) {
		if movieID == 1 {
			return &models.BudgetAndFees{ID: 10, IDMovie: 1, TotalBudget: 100}, nil
		}
		return nil, errBoom
	}}

	rec := serve("GET /movies/{id}/budget", GetBudget(store), http.MethodGet, "/movies/1/budget", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var b models.BudgetAndFees
	if err := json.NewDecoder(rec.Body).Decode(&b); err != nil || b.IDMovie != 1 || b.TotalBudget != 100 {
		t.Errorf("unexpected body: %+v (err %v)", b, err)
	}

	if rec := serve("GET /movies/{id}/budget", GetBudget(store), http.MethodGet, "/movies/2/budget", ""); rec.Code != http.StatusNotFound {
		t.Errorf("missing budget status = %d, want 404", rec.Code)
	}
	if rec := serve("GET /movies/{id}/budget", GetBudget(store), http.MethodGet, "/movies/x/budget", ""); rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}
}

func TestUpdateBudgetByMovie(t *testing.T) {
	store := &fakeBudgets{updateFn: func(movieID int, total, prod, other float64) (*models.BudgetAndFees, error) {
		return &models.BudgetAndFees{IDMovie: movieID, TotalBudget: total}, nil
	}}
	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{"success", "/movies/2/budget", `{"total_budget":50}`, http.StatusOK},
		{"invalid id", "/movies/x/budget", `{}`, http.StatusBadRequest},
		{"invalid body", "/movies/2/budget", `[`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("PUT /movies/{id}/budget", UpdateBudgetByMovie(store), http.MethodPut, tt.target, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestDeleteBudget(t *testing.T) {
	var deleted int
	store := &fakeBudgets{deleteFn: func(id int) error { deleted = id; return nil }}

	rec := serve("DELETE /budget/{id}", DeleteBudget(store), http.MethodDelete, "/budget/12", "")
	if rec.Code != http.StatusNoContent || deleted != 12 {
		t.Errorf("status = %d, deleted = %d; want 204 and 12", rec.Code, deleted)
	}
	if rec := serve("DELETE /budget/{id}", DeleteBudget(store), http.MethodDelete, "/budget/x", ""); rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}
	failing := &fakeBudgets{deleteFn: func(int) error { return errBoom }}
	if rec := serve("DELETE /budget/{id}", DeleteBudget(failing), http.MethodDelete, "/budget/12", ""); rec.Code != http.StatusInternalServerError {
		t.Errorf("storage failure status = %d, want 500", rec.Code)
	}
}
