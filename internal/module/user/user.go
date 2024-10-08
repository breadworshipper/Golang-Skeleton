package user

import (
	handler "mm-pddikti-cms/internal/module/user/handler/rest"
	"mm-pddikti-cms/internal/module/user/repository"
	"mm-pddikti-cms/internal/module/user/service"
)

func Init() *handler.UserHandler {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)

	return handler.NewUserHandler(userService)
}
