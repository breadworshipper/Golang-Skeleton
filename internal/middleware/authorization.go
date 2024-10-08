package middleware

import (
	"mm-pddikti-cms/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthRole(authorizedRoles []string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
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
			Role           string   `json:"role"`
			AuthorizedRole []string `json:"authorized_roles"`
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
