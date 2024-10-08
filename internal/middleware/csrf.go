package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

func Csrf() func(*fiber.Ctx) error {
	sessionConfig := session.Config{
		// KeyLookup:      "cookie:" + config.AppConfig.Application.Session.Key, // Recommended to use the __Host- prefix when serving the app over TLS
		KeyLookup:      "cookie:session_csrf", // Recommended to use the __Host- prefix when serving the app over TLS
		Expiration:     time.Duration(120) * time.Minute,
		CookieSecure:   false,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}

	csrfConfig := csrf.Config{
		KeyLookup:         "header:X-Csrf-Token",
		CookieName:        "csrf_",
		CookieSameSite:    "Lax",
		CookieSecure:      false,
		ContextKey:        "csrf",
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
		Expiration:        time.Duration(1) * time.Minute,
		KeyGenerator:      utils.UUIDv4,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "CSRF Token not valid!",
				"success": false,
			})
		},
		Extractor:         csrf.CsrfFromHeader("X-Csrf-Token"),
		Session:           session.New(sessionConfig),
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
		// CookieName:        config.AppConfig.Application.CSRF.Key,
		// CookieSecure:      config.AppConfig.Application.CSRF.Secure,
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return httphelper.Response(
		// 		c, httphelper.ResponseParams{
		// 			StatusCode: fiber.StatusForbidden,
		// 			Message:    "CSRF Token not valid!",
		// 		})
		// },
	}

	return csrf.New(csrfConfig)
}
