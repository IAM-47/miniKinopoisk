package storage

import (
	"context"
	"database/sql"
	"fmt"
	"miniKinopoisk/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ActorStorage struct {
	db *pgxpool.Pool
}

func NewActorStorage(db *pgxpool.Pool) *ActorStorage {
	return &ActorStorage{db: db}
}

func (s *ActorStorage) CreateActor(ctx context.Context, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error) {
	query := `
		insert into actors(first_name, last_name, birth_date, salary)
		values ($1, $2, $3, $4)
		returning id, first_name, last_name, birth_date, salary;
	`

	var actor models.Actor
	var bd sql.NullTime

	if birthDate != nil {
		bd = sql.NullTime{Time: *birthDate, Valid: true}
	} else {
		bd = sql.NullTime{Valid: false}
	}

	err := s.db.QueryRow(ctx, query, firstName, lastName, bd, salary).Scan(
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
		select a.id, a.first_name, a.last_name, a.birth_date, a.salary from actors a 
		join movie_actor ma 
		on a.id = ma.id_actor 
		where ma.id_movie = $1;
		`
	rows, err := s.db.Query(ctx, query, movieID)
	if err != nil {
		return nil, fmt.Errorf("failed to get actors: %w", err)
	}
	defer rows.Close()

	var actors []*models.Actor
	for rows.Next() {
		var actor models.Actor
		var bd sql.NullTime
		if err := rows.Scan(
			&actor.ID,
			&actor.FirstName,
			&actor.LastName,
			&bd,
			&actor.Salary); err != nil {
			return nil, fmt.Errorf("failed to scan actors: %w", err)
		}
		if bd.Valid {
			actor.BirthDate = bd.Time
		}
		actors = append(actors, &actor)
	}
	return actors, nil
}

func (s *ActorStorage) AddActorToMovie(ctx context.Context, movieID, actorID int) error {
	query := `insert into movie_actor (id_movie, id_actor)
		values ($1, $2) on conflict (id_movie, id_actor) do nothing;`
	_, err := s.db.Exec(ctx, query, movieID, actorID)
	if err != nil {
		return fmt.Errorf("failed to add actor to movie: %w", err)
	}
	return nil
}

func (s *ActorStorage) UpdateActor(ctx context.Context, id int, firstName, lastName string, birthDate *time.Time, salary float64) (*models.Actor, error) {
	query := `
		update actors
		set first_name = $2, last_name = $3, birth_date = $4, salary = $5
		where id = $1
		returning id, first_name, last_name, birth_date, salary;
	`
	var actor models.Actor
	var bd sql.NullTime
	err := s.db.QueryRow(ctx, query, id, firstName, lastName, birthDate, salary).Scan(
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
	} else {
		actor.BirthDate = time.Time{}
	}
	return &actor, nil
}

func (s *ActorStorage) DeleteActor(ctx context.Context, id int) error {
	query := `delete from actors where id = $1;`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete actor: %w", err)
	}
	return nil
}
