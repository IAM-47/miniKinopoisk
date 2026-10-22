package utils

import "testing"

func TestHashPasswordAndCheck(t *testing.T) {
	const password = "s3cret-Passw0rd"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword error: %v", err)
	}
	if hash == password {
		t.Fatal("hash must not equal the plain password")
	}
	if !CheckPasswordHash(password, hash) {
		t.Error("correct password must match its hash")
	}
	if CheckPasswordHash("wrong-password", hash) {
		t.Error("wrong password must not match")
	}
}

func TestHashPasswordUsesSalt(t *testing.T) {
	h1, err1 := HashPassword("same-password")
	h2, err2 := HashPassword("same-password")
	if err1 != nil || err2 != nil {
		t.Fatalf("HashPassword errors: %v, %v", err1, err2)
	}
	if h1 == h2 {
		t.Error("bcrypt must produce different hashes for the same password (random salt)")
	}
}

func TestCheckPasswordHashInvalidHash(t *testing.T) {
	if CheckPasswordHash("password", "not-a-bcrypt-hash") {
		t.Error("garbage hash must never match")
	}
	if CheckPasswordHash("password", "") {
		t.Error("empty hash must never match")
	}
}
