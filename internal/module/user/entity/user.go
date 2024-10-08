package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleSuperAdmin Role = "super-admin"
	RoleAdmin      Role = "admin"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	FullName  string
	Username  string    `gorm:"unique"`
	Email     string    `gorm:"unique"`
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}
