package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/auth"
	"miniKinopoisk/internal/storage"
	"miniKinopoisk/internal/utils"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Email == "" || req.Password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		if !strings.Contains(req.Email, "@") {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}

		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Printf("HashPassword error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = userStorage.CreateUser(r.Context(), req.Email, hash)
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
				http.Error(w, "User with this email already exists", http.StatusConflict)
				return
			}
			log.Printf("CreateUser error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}

func Login(userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Email == "" || req.Password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		user, err := userStorage.GetUserByEmail(r.Context(), req.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
		if err != nil {
			log.Printf("GenerateToken error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func DeleteUserByEmail(userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")
		if email == "" {
			http.Error(w, "I`m waiting for email", http.StatusBadRequest)
			return
		}

		if err := userStorage.DeleteUserByEmail(r.Context(), email); err != nil {
			log.Printf("DeleteMovie error: %v", err)
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
