package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"miniKinopoisk/internal/models"
)

type ActorStorage struct {
	db *pgxpool.Pool
}

func NewActorStorage(db *pgxpool.Pool) *ActorStorage {
	return &ActorStorage{db: db}
}

func timeToPgDate(t *time.Time) pgtype.Date {
	if t == nil {
		return pgtype.Date{Valid: false}
	}
	return pgtype.Date{Time: *t, Valid: true}
}

func (s *ActorStorage) CreateActor(ctx context.Context, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error) {
	query := `
		INSERT INTO actors(first_name, last_name, birth_date, salary)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, birth_date, salary;
	`
	var actor models.Actor
	var bd pgtype.Date

	err := s.db.QueryRow(ctx, query, firstName, lastName, timeToPgDate(birthDate), salary).Scan(
		&actor.ID,
		&actor.FirstName,
		&actor.LastName,
		&bd,
		&actor.Salary,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create actor: %w", err)
	}
	if bd.Valid {
		actor.BirthDate = bd.Time
	}
	return &actor, nil
}

func (s *ActorStorage) GetActorsByMovie(ctx context.Context, movieID int) ([]*models.Actor, error) {
	query := `
		SELECT a.id, a.first_name, a.last_name, a.birth_date, a.salary
		FROM actors a
		JOIN movie_actor ma ON a.id = ma.id_actor
		WHERE ma.id_movie = $1;
	`
	rows, err := s.db.Query(ctx, query, movieID)
	if err != nil {
		return nil, fmt.Errorf("failed to get actors: %w", err)
	}
	defer rows.Close()

	actors := make([]*models.Actor, 0)
	for rows.Next() {
		var actor models.Actor
		var bd pgtype.Date
		if err := rows.Scan(
			&actor.ID,
			&actor.FirstName,
			&actor.LastName,
			&bd,
			&actor.Salary,
		); err != nil {
			return nil, fmt.Errorf("failed to scan actor: %w", err)
		}
		if bd.Valid {
			actor.BirthDate = bd.Time
		}
		actors = append(actors, &actor)
	}
	return actors, nil
}

func (s *ActorStorage) AddActorToMovie(ctx context.Context, movieID, actorID int) error {
	query := `
		INSERT INTO movie_actor (id_movie, id_actor)
		VALUES ($1, $2)
		ON CONFLICT (id_movie, id_actor) DO NOTHING;
	`
	_, err := s.db.Exec(ctx, query, movieID, actorID)
	if err != nil {
		return fmt.Errorf("failed to add actor to movie: %w", err)
	}
	return nil
}

func (s *ActorStorage) UpdateActor(ctx context.Context, id int, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error) {
	query := `
		UPDATE actors
		SET first_name = $2, last_name = $3, birth_date = $4, salary = $5
		WHERE id = $1
		RETURNING id, first_name, last_name, birth_date, salary;
	`
	var actor models.Actor
	var bd pgtype.Date
	err := s.db.QueryRow(ctx, query, id, firstName, lastName, timeToPgDate(birthDate), salary).Scan(
		&actor.ID,
		&actor.FirstName,
		&actor.LastName,
		&bd,
		&actor.Salary,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update actor: %w", err)
	}
	if bd.Valid {
		actor.BirthDate = bd.Time
	}
	return &actor, nil
}

func (s *ActorStorage) DeleteActor(ctx context.Context, id int) error {
	query := `DELETE FROM actors WHERE id = $1;`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete actor: %w", err)
	}
	return nil
}
