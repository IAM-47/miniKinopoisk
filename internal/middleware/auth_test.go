package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"miniKinopoisk/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "middleware-test-secret"

func signToken(t *testing.T, method jwt.SigningMethod, secret string, role string, ttl time.Duration) string {
	t.Helper()
	claims := &auth.Claims{
		UserID: 7,
		Email:  "u@example.com",
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}
	s, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("sign token: %v", err)
	}
	return s
}

func TestAuthMiddleware(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	valid := signToken(t, jwt.SigningMethodHS256, testSecret, "user", time.Hour)
	expired := signToken(t, jwt.SigningMethodHS256, testSecret, "user", -time.Hour)
	wrongSecret := signToken(t, jwt.SigningMethodHS256, "another-secret", "user", time.Hour)

	// Токен с alg=none: классическая атака на JWT, должен отклоняться.
	noneClaims := &auth.Claims{UserID: 1, Role: "admin"}
	noneToken, err := jwt.NewWithClaims(jwt.SigningMethodNone, noneClaims).SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("sign none token: %v", err)
	}

	tests := []struct {
		name       string
		header     string
		wantStatus int
	}{
		{"no header", "", http.StatusUnauthorized},
		{"not a bearer scheme", "Basic abc123", http.StatusUnauthorized},
		{"garbage token", "Bearer not.a.jwt", http.StatusUnauthorized},
		{"expired token", "Bearer " + expired, http.StatusUnauthorized},
		{"signed with wrong secret", "Bearer " + wrongSecret, http.StatusUnauthorized},
		{"alg none token", "Bearer " + noneToken, http.StatusUnauthorized},
		{"valid token", "Bearer " + valid, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			h := AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				called = true
			})

			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			if tt.header != "" {
				req.Header.Set("Authorization", tt.header)
			}
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if wantCalled := tt.wantStatus == http.StatusOK; called != wantCalled {
				t.Errorf("next handler called = %v, want %v", called, wantCalled)
			}
		})
	}
}

func TestAuthMiddlewarePutsClaimsIntoContext(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)
	token := signToken(t, jwt.SigningMethodHS256, testSecret, "admin", time.Hour)

	var gotID int
	var gotRole string
	h := AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		gotID, _ = r.Context().Value(ContextKeyUserID).(int)
		gotRole, _ = r.Context().Value(ContextKeyRole).(string)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	h.ServeHTTP(httptest.NewRecorder(), req)

	if gotID != 7 || gotRole != "admin" {
		t.Errorf("context has user_id=%d role=%q, want 7 and admin", gotID, gotRole)
	}
}

// Полная цепочка, как она собрана в app.go: AuthMiddleware -> AdminOnly -> handler.
func TestAdminOnlyChain(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	chain := AuthMiddleware(AdminOnly(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	tests := []struct {
		name       string
		role       string
		wantStatus int
	}{
		{"admin gets access", "admin", http.StatusNoContent},
		{"regular user is forbidden", "user", http.StatusForbidden},
		{"empty role is forbidden", "", http.StatusForbidden},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/movies/1", nil)
			req.Header.Set("Authorization", "Bearer "+signToken(t, jwt.SigningMethodHS256, testSecret, tt.role, time.Hour))
			rec := httptest.NewRecorder()
			chain.ServeHTTP(rec, req)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestAdminOnlyWithoutAuthContext(t *testing.T) {
	h := AdminOnly(func(w http.ResponseWriter, r *http.Request) {
		t.Error("handler must not be called without a role in context")
	})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusForbidden {
		t.Errorf("status = %d, want 403", rec.Code)
	}
}
