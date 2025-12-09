package storage

import (
	"context"
	"miniKinopoisk/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(ctx context.Context, email, hash string) (*models.User, error) {
	query := `
		insert into users (email, password_hash)
		values ($1, $2)
		returning id, email, password_hash, role, created_at;
	`
	var u models.User
	err := s.db.QueryRow(ctx, query, email, hash).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
	)
	return &u, err
}
