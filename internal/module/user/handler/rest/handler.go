package handler

import (
	"pddikti-cms/internal/module/user/ports"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service ports.UserService
}

func NewUserHandler(s ports.UserService) *userHandler {
	return &userHandler{
		service: s,
	}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	return nil
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	return nil
}
