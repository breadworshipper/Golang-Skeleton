package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	FullName  string     `json:"full_name" validate:"required"`
	Username  string     `json:"username" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"password" validate:"required"`
	Role      string     `json:"role" validate:"required,oneof=super-admin admin"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
