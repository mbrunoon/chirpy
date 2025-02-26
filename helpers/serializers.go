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
