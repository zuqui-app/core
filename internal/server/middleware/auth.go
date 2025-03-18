package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"zuqui/internal/service/auth"
)

func AuthMiddleware(as auth.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		parts := strings.Split(c.Get("Authorization"), "Bearer ")
		if len(parts) < 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := parts[1]
		_, err := as.VerifyAccessToken(token)
		if err != nil {
			if _, ok := err.(*auth.AuthError); ok {
				log.Println(err)
				return fiber.NewError(fiber.StatusUnauthorized)
			}
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.Next()
	}
}
