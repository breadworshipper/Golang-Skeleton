package route

import (
	handler "mm-pddikti-cms/internal/module/auth/handler/rest"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app fiber.Router, authHandler *handler.AuthHandler) {
	group := app.Group("/auth")
	group.Post("/login", authHandler.Login)
}
