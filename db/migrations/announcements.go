package migrations

import "time"

type Announcements struct {
	ID          string     `gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title       string     `gorm:"column:title"`
	Slug        string     `gorm:"column:slug"`
	Link        string     `gorm:"column:link"`
	Description string     `gorm:"column:description"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}
