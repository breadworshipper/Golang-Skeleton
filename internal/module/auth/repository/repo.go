package repository

import (
	"context"
	"mm-pddikti-cms/internal/adapter"
	"mm-pddikti-cms/internal/module/auth/ports"
	"mm-pddikti-cms/internal/module/user/entity"

	"gorm.io/gorm"
)

var _ ports.AuthRepository = (*authRepository)(nil)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() ports.AuthRepository {
	return &authRepository{
		db: adapter.Adapters.Postgres,
	}
}

func (r *authRepository) FindUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

