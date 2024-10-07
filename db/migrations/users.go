package migrations

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string     `gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	FullName  string     `gorm:"column:full_name"`
	Username  string     `gorm:"column:username"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	Role      string     `gorm:"column:role"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
