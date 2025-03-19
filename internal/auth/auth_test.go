package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{"Correct password", password1, hash1, false},
		{"Incorrect password", "wrongPassword", hash1, true},
		{"Password doesn't match different hash", password1, hash2, true},
		{"Empty password", "", hash1, true},
		{"Invalid hash", password1, "invalidhash", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("CheckPasswordHash() for %q: gotErr = %v, wantErr = %v", tt.name, gotErr, tt.wantErr)
			}
		})
	}
}

func TestMakeJWTAndValidateJWT(t *testing.T) {
	tokenSecret := "supersecretkey"
	userID := uuid.New()
	expiresIn := time.Minute * 5

	t.Run("Valid JWT", func(t *testing.T) {
		token, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("Failed to create JWT: %v", err)
		}
		if token == "" {
			t.Fatalf("Expected non-empty token")
		}

		parsedUserID, err := ValidateJWT(token, tokenSecret)
		if err != nil {
			t.Errorf("ValidateJWT() returned error: %v", err)
		}
		if parsedUserID != userID {
			t.Errorf("Expected userID %v, got %v", userID, parsedUserID)
		}
	})

	t.Run("Expired JWT", func(t *testing.T) {
		expiredToken, err := MakeJWT(userID, tokenSecret, -time.Minute) // Token expirado
		if err != nil {
			t.Fatalf("Failed to create expired JWT: %v", err)
		}

		_, err = ValidateJWT(expiredToken, tokenSecret)
		if err == nil {
			t.Errorf("Expected error for expired token, got nil")
		}
	})

	t.Run("Invalid Secret", func(t *testing.T) {
		token, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("Failed to create JWT: %v", err)
		}

		_, err = ValidateJWT(token, "wrongsecret")
		if err == nil {
			t.Errorf("Expected error for wrong secret, got nil")
		}
	})

	t.Run("Invalid Token Format", func(t *testing.T) {
		_, err := ValidateJWT("invalid.token.string", tokenSecret)
		if err == nil {
			t.Errorf("Expected error for invalid token format, got nil")
		}
	})
}
