package storage

import (
	"context"
	"fmt"
	"miniKinopoisk/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetStorage struct {
	db *pgxpool.Pool
}

func NewBudgetStorage(db *pgxpool.Pool) *BudgetStorage {
	return &BudgetStorage{db: db}
}

func (s *BudgetStorage) CreateBudget(ctx context.Context, movieID int, totalBudget, feesInProdCountry, feesinOther float64) (*models.BudgetAndFees, error) {
	query := `
		insert into budget_and_fees(movieID, total_budget, fees_in_prod_country, fees_in_other)
		values ($1, $2, $3, $4)
		on conflict (id_movie) do update
		set total_budget = $2, fees_in_prod_country = $3, fees_in_other = $4
		returning id, movieID, total_budget, fees_in_prod_country, fees_in_other;
	`

	var budget models.BudgetAndFees

	err := s.db.QueryRow(ctx, query, movieID, totalBudget, feesInProdCountry, feesinOther).Scan(
		&budget.ID,
		&budget.IDMovie,
		&budget.TotalBudget,
		&budget.FeesInProdCountry,
		&budget.FeesInOther,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create budget: %w", err)
	}

	return &budget, nil

}

func (s *BudgetStorage) GetBudgetByMovie(ctx context.Context, movieID int) (*models.BudgetAndFees, error) {
	query := `select * from budget_and_fees where id_movie = $1;`

	var budget models.BudgetAndFees
	err := s.db.QueryRow(ctx, query, movieID).Scan(
		&budget.ID,
		&budget.IDMovie,
		&budget.TotalBudget,
		&budget.FeesInProdCountry,
		&budget.FeesInOther,
	)
	if err != nil {
		return nil, fmt.Errorf("budget is not found for movie %d: %w", movieID, err)
	}
	return &budget, nil
}

func (s *BudgetStorage) UpdateBudgetByMovie(ctx context.Context, movieID int, totalBudget, feesInProdCountry, feeInOther float64) (*models.BudgetAndFees, error) {
	query := `
		update budget_and_fees
		set total_budget = $2, fees_in_prod_country = $3, fees_in_other = $4
		where id_movie = $1
		on conflict (id_movie) do update
		returning id, total_budget, fees_in_prod_country, fees_in_other;
	`
	var budget models.BudgetAndFees
	err := s.db.QueryRow(ctx, query, movieID, totalBudget, feesInProdCountry, feeInOther).Scan(
		&budget.ID,
		&budget.IDMovie,
		&budget.TotalBudget,
		&budget.FeesInProdCountry,
		&budget.FeesInOther,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update budget: %w", err)
	}
	return &budget, nil
}

func (s *BudgetStorage) DeleteBudget(ctx context.Context, id int) error {
	query := `delete from budget_and_fees where id = $1;`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete budget: %w", err)
	}
	return nil
}
