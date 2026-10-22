package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"miniKinopoisk/internal/models"
)

func TestCreateMovie(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		storeErr   error
		wantStatus int
	}{
		{"success", `{"title":"Dune","producer":"Mary","director":"Denis","release_year":2021}`, nil, http.StatusCreated},
		{"invalid json", `nope`, nil, http.StatusBadRequest},
		{"missing title", `{"producer":"Mary","director":"Denis"}`, nil, http.StatusBadRequest},
		{"missing producer", `{"title":"Dune","director":"Denis"}`, nil, http.StatusBadRequest},
		{"missing director", `{"title":"Dune","producer":"Mary"}`, nil, http.StatusBadRequest},
		{"storage failure", `{"title":"Dune","producer":"Mary","director":"Denis"}`, errBoom, http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeMovies{createFn: func(title, producer, director string, year int) (*models.Movie, error) {
				if tt.storeErr != nil {
					return nil, tt.storeErr
				}
				return &models.Movie{ID: 1, Title: title, Producer: producer, Director: director, ReleaseYear: year}, nil
			}}
			rec := serve("POST /movies", CreateMovie(store), http.MethodPost, "/movies", tt.body)
			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantStatus == http.StatusCreated {
				var m models.Movie
				if err := json.NewDecoder(rec.Body).Decode(&m); err != nil {
					t.Fatalf("invalid JSON: %v", err)
				}
				if m.Title != "Dune" || m.ReleaseYear != 2021 {
					t.Errorf("unexpected movie in response: %+v", m)
				}
			}
		})
	}
}

func TestGetMoviesPagination(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantLimit  int
		wantOffset int
	}{
		{"defaults", "/movies", 10, 0},
		{"custom values", "/movies?limit=5&offset=20", 5, 20},
		{"zero limit falls back to default", "/movies?limit=0", 10, 0},
		{"negative limit falls back to default", "/movies?limit=-3", 10, 0},
		{"negative offset falls back to zero", "/movies?offset=-1", 10, 0},
		{"garbage values fall back", "/movies?limit=abc&offset=xyz", 10, 0},
		{"zero offset is allowed", "/movies?limit=2&offset=0", 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotLimit, gotOffset int
			store := &fakeMovies{listFn: func(limit, offset int) ([]*models.Movie, error) {
				gotLimit, gotOffset = limit, offset
				return []*models.Movie{{ID: 1, Title: "Dune"}}, nil
			}}
			rec := serve("GET /movies", GetMovies(store), http.MethodGet, tt.query, "")
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want 200", rec.Code)
			}
			if gotLimit != tt.wantLimit || gotOffset != tt.wantOffset {
				t.Errorf("store got limit=%d offset=%d, want limit=%d offset=%d", gotLimit, gotOffset, tt.wantLimit, tt.wantOffset)
			}
		})
	}
}

func TestGetMoviesStorageFailure(t *testing.T) {
	store := &fakeMovies{listFn: func(int, int) ([]*models.Movie, error) { return nil, errBoom }}
	rec := serve("GET /movies", GetMovies(store), http.MethodGet, "/movies", "")
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", rec.Code)
	}
}

func TestGetMovieByID(t *testing.T) {
	store := &fakeMovies{getFn: func(id int) (*models.Movie, error) {
		if id == 1 {
			return &models.Movie{ID: 1, Title: "Dune"}, nil
		}
		return nil, errBoom
	}}

	tests := []struct {
		target     string
		wantStatus int
	}{
		{"/movies/1", http.StatusOK},
		{"/movies/999", http.StatusNotFound},
		{"/movies/abc", http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			rec := serve("GET /movies/{id}", GetMovieByID(store), http.MethodGet, tt.target, "")
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	var gotID int
	store := &fakeMovies{updateFn: func(id int, title, producer, director string, year int) (*models.Movie, error) {
		gotID = id
		return &models.Movie{ID: id, Title: title, Producer: producer, Director: director, ReleaseYear: year}, nil
	}}
	valid := `{"title":"Dune 2","producer":"Mary","director":"Denis","release_year":2024}`

	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{"success", "/movies/3", valid, http.StatusOK},
		{"invalid id", "/movies/x", valid, http.StatusBadRequest},
		{"invalid body", "/movies/3", `{`, http.StatusBadRequest},
		{"missing required field", "/movies/3", `{"title":"Dune 2"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("PUT /movies/{id}", UpdateMovie(store), http.MethodPut, tt.target, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
	if gotID != 3 {
		t.Errorf("id passed to storage = %d, want 3", gotID)
	}
}

func TestDeleteMovie(t *testing.T) {
	var deleted int
	store := &fakeMovies{deleteFn: func(id int) error { deleted = id; return nil }}

	rec := serve("DELETE /movies/{id}", DeleteMovie(store), http.MethodDelete, "/movies/4", "")
	if rec.Code != http.StatusNoContent || deleted != 4 {
		t.Errorf("status = %d, deleted id = %d; want 204 and 4", rec.Code, deleted)
	}

	rec = serve("DELETE /movies/{id}", DeleteMovie(store), http.MethodDelete, "/movies/abc", "")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}

	failing := &fakeMovies{deleteFn: func(int) error { return errBoom }}
	rec = serve("DELETE /movies/{id}", DeleteMovie(failing), http.MethodDelete, "/movies/4", "")
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("storage failure status = %d, want 500", rec.Code)
	}
}

func TestGetMoviesByActor(t *testing.T) {
	var gotActor int
	store := &fakeMovies{byActorFn: func(actorID int) ([]*models.Movie, error) {
		gotActor = actorID
		return []*models.Movie{{ID: 1, Title: "Dune"}}, nil
	}}

	rec := serve("GET /actors/{id}/movies", GetMoviesByActor(store), http.MethodGet, "/actors/8/movies", "")
	if rec.Code != http.StatusOK || gotActor != 8 {
		t.Errorf("status = %d, actor id = %d; want 200 and 8", rec.Code, gotActor)
	}

	rec = serve("GET /actors/{id}/movies", GetMoviesByActor(store), http.MethodGet, "/actors/abc/movies", "")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}
}
