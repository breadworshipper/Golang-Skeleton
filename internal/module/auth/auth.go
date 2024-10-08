package auth

import (
	handler "mm-pddikti-cms/internal/module/auth/handler/rest"
	"mm-pddikti-cms/internal/module/auth/repository"
	"mm-pddikti-cms/internal/module/auth/service"
)

func Init() *handler.AuthHandler {
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository)

	return handler.NewAuthHandler(authService)
}
