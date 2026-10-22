package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func parse(t *testing.T, tokenStr, secret string) (*Claims, error) {
	t.Helper()
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return claims, err
}

func TestGenerateTokenContainsClaims(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")

	tokenStr, err := GenerateToken(42, "user@example.com", "admin")
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	claims, err := parse(t, tokenStr, "test-secret")
	if err != nil {
		t.Fatalf("token must be valid with the same secret: %v", err)
	}
	if claims.UserID != 42 || claims.Email != "user@example.com" || claims.Role != "admin" {
		t.Errorf("unexpected claims: %+v", claims)
	}
}

func TestGenerateTokenExpiresInAboutADay(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")

	tokenStr, err := GenerateToken(1, "a@b.c", "user")
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	claims, err := parse(t, tokenStr, "test-secret")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	left := time.Until(claims.ExpiresAt.Time)
	if left < 23*time.Hour || left > 24*time.Hour+time.Minute {
		t.Errorf("token lifetime should be ~24h, got %v", left)
	}
}

func TestTokenRejectedWithWrongSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", "right-secret")

	tokenStr, err := GenerateToken(1, "a@b.c", "user")
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	if _, err := parse(t, tokenStr, "other-secret"); err == nil {
		t.Error("token signed with another secret must be rejected")
	}
}

func TestJWTSecretIsReadFromEnv(t *testing.T) {
	t.Setenv("JWT_SECRET", "from-env")
	if got := string(jwtKey()); got != "from-env" {
		t.Errorf("jwtKey() = %q, want value from JWT_SECRET", got)
	}
}
