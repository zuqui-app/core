package server

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/service/auth"
)

func (a *App) AuthMiddleware(c *fiber.Ctx) error {
	parts := strings.Split(c.Get("Authorization"), "Bearer ")
	if len(parts) < 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := parts[1]
	_, err := a.auth.VerifyAccessToken(token)
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
