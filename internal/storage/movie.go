package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"miniKinopoisk/internal/models"
)

type MovieStorage struct {
	db *pgxpool.Pool
}

func NewMovieStorage(db *pgxpool.Pool) *MovieStorage {
	return &MovieStorage{db: db}
}

func (s *MovieStorage) CreateMovie(ctx context.Context, title, producer, director string, releaseYear int) (*models.Movie, error) {
	query := `
		insert into movies (title, producer, director, release_year)
		values ($1, $2, $3, $4)
		returning id, title, producer, director, release_year;
	`
	var movie models.Movie
	err := s.db.QueryRow(ctx, query, title, producer, director, releaseYear).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Producer,
		&movie.Director,
		&movie.ReleaseYear,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create movie: %w", err)
	}
	return &movie, nil
}

func (s *MovieStorage) GetMovies(ctx context.Context) ([]*models.Movie, error) {
	query := `select * from movies;`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get movies: %w", err)
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Producer,
			&movie.Director,
			&movie.ReleaseYear,
		); err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (s *MovieStorage) UpdateMovie(ctx context.Context, id int, title, producer, director string, release_year int) (*models.Movie, error) {
	query := `
		update movies
		set title = $2, producer = $3, director = $4, release_year = $5
		where id = $1
		returning id, title, producer, director, release_year;
	`
	var movie models.Movie
	err := s.db.QueryRow(ctx, query, id, title, producer, director, release_year).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Producer,
		&movie.Director,
		&movie.ReleaseYear,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update movie: %w", err)
	}
	return &movie, nil
}

func (s *MovieStorage) DeleteMovie(ctx context.Context, id int) error {
	query := `delete from movies where id = $1;`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %w", err)
	}
	return nil
}
