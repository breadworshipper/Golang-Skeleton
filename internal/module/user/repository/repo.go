package repository

import (
	"context"
	"mm-pddikti-cms/internal/adapter"
	"mm-pddikti-cms/internal/module/user/entity"
	"mm-pddikti-cms/internal/module/user/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ports.UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() ports.UserRepository {
	return &userRepository{
		db: adapter.Adapters.Postgres,
	}
}

func (r *userRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) IsFieldUnique(ctx context.Context, field, value string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(field+" = ?", value).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}
