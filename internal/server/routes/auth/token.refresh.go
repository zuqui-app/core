package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui/internal/service/auth"
)

type TokenRefreshRequest struct {
	Refresh string `json:"refresh" xml:"refresh" form:"refresh"`
}

type TokenRefreshResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// Refresh Tokens
func AuthTokenRefresh(as auth.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req TokenRefreshRequest

		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "malformed request")
		}

		if req.Refresh == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token is required")
		}

		claims, err := as.VerifyRefreshToken(req.Refresh)
		if err != nil {
			if err, ok := err.(*auth.AuthError); ok {
				log.Println(err)
				return fiber.NewError(fiber.StatusUnauthorized)
			}
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		access, refresh, err := as.CreateTokenPair(claims.Subject)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		if err = as.RevokeRefreshToken(claims.ID); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.JSON(TokenRefreshResponse{
			Access:  access,
			Refresh: refresh,
		})
	}
}
