package repository

import (
	"pddikti-cms/internal/module/user/entity"
	"pddikti-cms/internal/module/user/ports"

	"gorm.io/gorm"
)

var _ ports.UserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) EmailExist(email string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) Register(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
