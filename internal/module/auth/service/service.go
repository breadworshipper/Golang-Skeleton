package service

import (
	"context"
	"errors"
	"mm-pddikti-cms/internal/infrastructure/config"
	"mm-pddikti-cms/internal/module/auth/ports"
	"mm-pddikti-cms/pkg"
	"mm-pddikti-cms/pkg/jwthandler"
	"time"
)

var _ ports.AuthService = (*authService)(nil)

type authService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) ports.AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(context context.Context, usernameOrEmail, password string) (string, string, error) {
	user, err := s.repo.FindUserByUsernameOrEmail(context, usernameOrEmail)
	if err != nil {
		return "", "", errors.New("Invalid credentials!")
	}

	correct := pkg.ComparePassword(user.Password, password)
	if !correct {
		return "", "", errors.New("Invalid credentials!")
	}

	accessTokenString, refreshTokenString, err := jwthandler.GenerateTokenPairString(jwthandler.CostumClaimsPayload{
		UserId:                 user.ID,
		Username:               user.Username,
		Email:                  user.Email,
		Role:                   user.Role,
		AccessTokenExpiration:  time.Now().Add(time.Hour * time.Duration(config.Envs.Guard.JwtAccessTokenExpiration)),
		RefreshTokenExpiration: time.Now().Add(time.Hour * time.Duration(config.Envs.Guard.JwtRefreshTokenExpiration)),
	})
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}
