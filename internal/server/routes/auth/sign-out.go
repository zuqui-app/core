package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui/internal/service/auth"
)

type SignOutRequest struct {
	Refresh string `json:"refresh" xml:"refresh" form:"refresh"`
}

func SignOut(as auth.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req TokenRefreshRequest

		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "malformed request")
		}

		if req.Refresh == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token is required")
		}

		if err := as.RevokeRefreshToken(req.Refresh); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
