package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSSetsHeaders(t *testing.T) {
	h := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusTeapot {
		t.Errorf("regular request must reach the next handler, got status %d", rec.Code)
	}
	for header, want := range map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	} {
		if got := rec.Header().Get(header); got != want {
			t.Errorf("%s = %q, want %q", header, got, want)
		}
	}
}

func TestCORSPreflightDoesNotCallNext(t *testing.T) {
	called := false
	h := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodOptions, "/movies", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("preflight status = %d, want 200", rec.Code)
	}
	if called {
		t.Error("preflight (OPTIONS) must be answered by the middleware itself")
	}
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("preflight response must carry CORS headers")
	}
}
