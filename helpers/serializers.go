package helpers

import (
	"time"

	"github.com/google/uuid"
	"github.com/mbrunoon/chirpy/app/models"
)

type UserSerialized struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func UserSerializer(user *models.User) UserSerialized {
	return UserSerialized{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
}

type ChirpySerialized struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func ChirpySerializer(chirpy *models.Chirp) ChirpySerialized {
	return ChirpySerialized{
		ID:        chirpy.ID,
		CreatedAt: chirpy.CreatedAt,
		UpdatedAt: chirpy.UpdatedAt,
		Body:      chirpy.Body,
		UserID:    chirpy.UserID,
	}
}

func ChirpySerializerList(chirpies *[]models.Chirp) []ChirpySerialized {
	var chirpiesSerialized []ChirpySerialized
	for _, chirpy := range *chirpies {
		chirpiesSerialized = append(chirpiesSerialized, ChirpySerializer(&chirpy))
	}
	return chirpiesSerialized
}
