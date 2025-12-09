package handlers

import (
	"encoding/json"
	"log"
	"miniKinopoisk/internal/storage"
	"miniKinopoisk/internal/utils"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Email == "" || req.Password == "" {
			http.Error(w, "email and password are required", http.StatusBadRequest)
			return
		}

		//пока что сохраняем без хеша
		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = userStorage.CreateUser(r.Context(), req.Email, hash)
		if err != nil {
			if err.Error() == `ERROR: duplicate key value violates unique constraint "user_email_key" (SQLSTATE 23505)` {
				http.Error(w, "User with email already exists", http.StatusConflict)
				return
			}
			log.Printf("Registration error: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"meessage":"User created successfully!"}`))
	}
}

func Login(userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Logged in successfully",
			"user": map[string]interface{}{
				"id":    user.ID,
				"email": user.Email,
				"role":  user.Role,
			},
		})
	}
}
