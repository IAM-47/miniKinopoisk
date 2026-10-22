package handlers

import (
	"net/http"
	"testing"
	"time"

	"miniKinopoisk/internal/models"
)

func TestParseBirthDate(t *testing.T) {
	if got, err := parseBirthDate(nil); got != nil || err != nil {
		t.Errorf("nil input: got (%v, %v), want (nil, nil)", got, err)
	}

	valid := "1990-05-17"
	got, err := parseBirthDate(&valid)
	if err != nil {
		t.Fatalf("valid date returned error: %v", err)
	}
	if want := time.Date(1990, 5, 17, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
		t.Errorf("parsed %v, want %v", got, want)
	}

	for _, bad := range []string{"17.05.1990", "1990/05/17", "1990-13-01", ""} {
		bad := bad
		if _, err := parseBirthDate(&bad); err == nil {
			t.Errorf("date %q must be rejected", bad)
		}
	}
}

func TestCreateActor(t *testing.T) {
	var gotBirth *time.Time
	store := &fakeActors{createFn: func(first, last string, birth *time.Time, salary float64) (*models.Actor, error) {
		gotBirth = birth
		return &models.Actor{ID: 1, FirstName: first, LastName: last, BirthDate: birth, Salary: salary}, nil
	}}

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"success with date", `{"first_name":"Timothee","last_name":"Chalamet","birth_date":"1995-12-27","salary":1000.5}`, http.StatusCreated},
		{"success without date", `{"first_name":"Timothee","last_name":"Chalamet"}`, http.StatusCreated},
		{"invalid json", `{`, http.StatusBadRequest},
		{"missing last name", `{"first_name":"Timothee"}`, http.StatusBadRequest},
		{"missing first name", `{"last_name":"Chalamet"}`, http.StatusBadRequest},
		{"bad date format", `{"first_name":"T","last_name":"C","birth_date":"27-12-1995"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("POST /actors", CreateActor(store), http.MethodPost, "/actors", tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}

	// последний успешный вызов был без даты -> в хранилище должен уйти nil
	serve("POST /actors", CreateActor(store), http.MethodPost, "/actors", `{"first_name":"A","last_name":"B","birth_date":"2000-01-02"}`)
	if gotBirth == nil || gotBirth.Format("2006-01-02") != "2000-01-02" {
		t.Errorf("birth date passed to storage = %v, want 2000-01-02", gotBirth)
	}
}

func TestCreateActorStorageFailure(t *testing.T) {
	store := &fakeActors{createFn: func(string, string, *time.Time, float64) (*models.Actor, error) { return nil, errBoom }}
	rec := serve("POST /actors", CreateActor(store), http.MethodPost, "/actors", `{"first_name":"A","last_name":"B"}`)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", rec.Code)
	}
}

func TestGetActorByID(t *testing.T) {
	store := &fakeActors{getFn: func(id int) (*models.Actor, error) {
		if id == 1 {
			return &models.Actor{ID: 1, FirstName: "A", LastName: "B"}, nil
		}
		return nil, errBoom
	}}
	tests := []struct {
		target     string
		wantStatus int
	}{
		{"/actors/1", http.StatusOK},
		{"/actors/2", http.StatusNotFound},
		{"/actors/x", http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			rec := serve("GET /actors/{id}", GetActorByID(store), http.MethodGet, tt.target, "")
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestGetActorsByMovie(t *testing.T) {
	var gotMovie int
	store := &fakeActors{byMovieFn: func(movieID int) ([]*models.Actor, error) {
		gotMovie = movieID
		return []*models.Actor{{ID: 1}}, nil
	}}
	rec := serve("GET /movies/{id}/actors", GetActorsByMovie(store), http.MethodGet, "/movies/6/actors", "")
	if rec.Code != http.StatusOK || gotMovie != 6 {
		t.Errorf("status = %d, movie id = %d; want 200 and 6", rec.Code, gotMovie)
	}
	rec = serve("GET /movies/{id}/actors", GetActorsByMovie(store), http.MethodGet, "/movies/q/actors", "")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}
}

func TestAddActorToMovie(t *testing.T) {
	var gotMovie, gotActor int
	store := &fakeActors{linkFn: func(movieID, actorID int) error {
		gotMovie, gotActor = movieID, actorID
		return nil
	}}

	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{"success", "/movies/2/actors", `{"actor_id":9}`, http.StatusCreated},
		{"invalid movie id", "/movies/z/actors", `{"actor_id":9}`, http.StatusBadRequest},
		{"invalid body", "/movies/2/actors", `{`, http.StatusBadRequest},
		{"zero actor id", "/movies/2/actors", `{"actor_id":0}`, http.StatusBadRequest},
		{"negative actor id", "/movies/2/actors", `{"actor_id":-4}`, http.StatusBadRequest},
		{"missing actor id", "/movies/2/actors", `{}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("POST /movies/{id}/actors", AddActorToMovie(store), http.MethodPost, tt.target, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
	if gotMovie != 2 || gotActor != 9 {
		t.Errorf("link called with movie=%d actor=%d, want 2 and 9", gotMovie, gotActor)
	}

	failing := &fakeActors{linkFn: func(int, int) error { return errBoom }}
	rec := serve("POST /movies/{id}/actors", AddActorToMovie(failing), http.MethodPost, "/movies/2/actors", `{"actor_id":9}`)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("storage failure status = %d, want 500", rec.Code)
	}
}

func TestUpdateActor(t *testing.T) {
	store := &fakeActors{updateFn: func(id int, first, last string, birth *time.Time, salary float64) (*models.Actor, error) {
		return &models.Actor{ID: id, FirstName: first, LastName: last, BirthDate: birth, Salary: salary}, nil
	}}
	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{"success", "/actors/1", `{"first_name":"A","last_name":"B","birth_date":"1999-09-09"}`, http.StatusOK},
		{"invalid id", "/actors/x", `{"first_name":"A","last_name":"B"}`, http.StatusBadRequest},
		{"invalid body", "/actors/1", `{`, http.StatusBadRequest},
		{"missing names", "/actors/1", `{"first_name":"A"}`, http.StatusBadRequest},
		{"bad date", "/actors/1", `{"first_name":"A","last_name":"B","birth_date":"nope"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("PUT /actors/{id}", UpdateActor(store), http.MethodPut, tt.target, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestDeleteActor(t *testing.T) {
	var deleted int
	store := &fakeActors{deleteFn: func(id int) error { deleted = id; return nil }}

	rec := serve("DELETE /actors/{id}", DeleteActor(store), http.MethodDelete, "/actors/5", "")
	if rec.Code != http.StatusNoContent || deleted != 5 {
		t.Errorf("status = %d, deleted = %d; want 204 and 5", rec.Code, deleted)
	}
	rec = serve("DELETE /actors/{id}", DeleteActor(store), http.MethodDelete, "/actors/x", "")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("invalid id status = %d, want 400", rec.Code)
	}
}
