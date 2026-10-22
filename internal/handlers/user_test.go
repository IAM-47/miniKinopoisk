package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"miniKinopoisk/internal/models"
	"miniKinopoisk/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestRegister(t *testing.T) {
	var gotEmail, gotHash string

	tests := []struct {
		name       string
		body       string
		createErr  error
		wantStatus int
	}{
		{"success", `{"email":"a@b.com","password":"pass123"}`, nil, http.StatusCreated},
		{"invalid json", `{oops`, nil, http.StatusBadRequest},
		{"missing password", `{"email":"a@b.com"}`, nil, http.StatusBadRequest},
		{"missing email", `{"password":"pass123"}`, nil, http.StatusBadRequest},
		{"email without at sign", `{"email":"ab.com","password":"pass123"}`, nil, http.StatusBadRequest},
		{"duplicate email", `{"email":"a@b.com","password":"pass123"}`, &pgconn.PgError{Code: "23505"}, http.StatusConflict},
		{"storage failure", `{"email":"a@b.com","password":"pass123"}`, errBoom, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmail, gotHash = "", ""
			store := &fakeUsers{createFn: func(_ context.Context, email, hash string) (*models.User, error) {
				gotEmail, gotHash = email, hash
				if tt.createErr != nil {
					return nil, tt.createErr
				}
				return &models.User{ID: 1, Email: email}, nil
			}}

			rec := serve("POST /register", Register(store), http.MethodPost, "/register", tt.body)
			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d (body: %s)", rec.Code, tt.wantStatus, rec.Body.String())
			}

			if tt.wantStatus == http.StatusCreated {
				if gotEmail != "a@b.com" {
					t.Errorf("stored email = %q", gotEmail)
				}
				if gotHash == "pass123" || !utils.CheckPasswordHash("pass123", gotHash) {
					t.Error("password must be stored as a valid bcrypt hash, not as plain text")
				}
			}
		})
	}
}

func TestLogin(t *testing.T) {
	t.Setenv("JWT_SECRET", "login-test-secret")

	hash, err := utils.HashPassword("pass123")
	if err != nil {
		t.Fatal(err)
	}
	store := &fakeUsers{getFn: func(_ context.Context, email string) (*models.User, error) {
		if email != "admin@b.com" {
			return nil, errBoom // «пользователь не найден»
		}
		return &models.User{ID: 5, Email: email, PasswordHash: hash, Role: "admin"}, nil
	}}

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"success", `{"email":"admin@b.com","password":"pass123"}`, http.StatusOK},
		{"wrong password", `{"email":"admin@b.com","password":"nope"}`, http.StatusUnauthorized},
		{"unknown user", `{"email":"ghost@b.com","password":"pass123"}`, http.StatusUnauthorized},
		{"missing fields", `{"email":"admin@b.com"}`, http.StatusBadRequest},
		{"invalid json", `[]`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve("POST /login", Login(store), http.MethodPost, "/login", tt.body)
			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantStatus != http.StatusOK {
				return
			}

			var resp map[string]string
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil || resp["token"] == "" {
				t.Fatalf("expected JSON with token, got %q (err %v)", rec.Body.String(), err)
			}
			parsed, err := jwt.Parse(resp["token"], func(*jwt.Token) (interface{}, error) {
				return []byte("login-test-secret"), nil
			})
			if err != nil || !parsed.Valid {
				t.Fatalf("issued token is not valid: %v", err)
			}
			if role := parsed.Claims.(jwt.MapClaims)["role"]; role != "admin" {
				t.Errorf("role in token = %v, want admin", role)
			}
		})
	}
}

func TestLoginDoesNotRevealWhichPartIsWrong(t *testing.T) {
	hash, _ := utils.HashPassword("pass123")
	store := &fakeUsers{getFn: func(_ context.Context, email string) (*models.User, error) {
		if email == "real@b.com" {
			return &models.User{Email: email, PasswordHash: hash}, nil
		}
		return nil, errBoom
	}}

	wrongPass := serve("POST /login", Login(store), http.MethodPost, "/login", `{"email":"real@b.com","password":"bad"}`)
	noUser := serve("POST /login", Login(store), http.MethodPost, "/login", `{"email":"fake@b.com","password":"bad"}`)

	if wrongPass.Body.String() != noUser.Body.String() {
		t.Errorf("error messages differ and allow user enumeration: %q vs %q", wrongPass.Body.String(), noUser.Body.String())
	}
}

func TestDeleteUserByEmail(t *testing.T) {
	var deleted string
	store := &fakeUsers{deleteFn: func(_ context.Context, email string) error {
		deleted = email
		return nil
	}}

	rec := serve("DELETE /users/{email}", DeleteUserByEmail(store), http.MethodDelete, "/users/a@b.com", "")
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}
	if deleted != "a@b.com" {
		t.Errorf("deleted email = %q", deleted)
	}

	failing := &fakeUsers{deleteFn: func(context.Context, string) error { return errBoom }}
	rec = serve("DELETE /users/{email}", DeleteUserByEmail(failing), http.MethodDelete, "/users/a@b.com", "")
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status on storage error = %d, want 500", rec.Code)
	}
}
