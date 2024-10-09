package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	ID        uuid.UUID  `json:"id"`
	FullName  string     `json:"full_name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Role      Role       `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func ToUserResponseDTO(user User) UserResponseDTO {
	return UserResponseDTO{
		ID:        user.ID,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
