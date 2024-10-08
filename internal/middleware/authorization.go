package middleware

import (
	"mm-pddikti-cms/internal/module/user/entity"
	"mm-pddikti-cms/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthRole(authorizedRoles []entity.Role) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(entity.Role)
		if !ok {
			return response.SendResponse(c, response.ResponseParams{
				StatusCode: fiber.StatusForbidden,
				Message:    "Terlarang: role anda tidak diizinkan untuk mengakses resource ini",
			})

		}

		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				return c.Next()
			}
		}

		payload := struct {
			Role           entity.Role   `json:"role"`
			AuthorizedRole []entity.Role `json:"authorized_roles"`
		}{
			Role:           role,
			AuthorizedRole: authorizedRoles,
		}

		log.Warn().Any("payload", payload).Msg("middleware::AuthRole - Unauthorized")
		return response.SendResponse(c, response.ResponseParams{
			StatusCode: fiber.StatusForbidden,
			Message:    "Terlarang: role anda tidak diizinkan untuk mengakses resource ini",
		})
	}
}
