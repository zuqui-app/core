package auth

import "github.com/gofiber/fiber/v2"

// Refresh Tokens
func TokenRefresh(c *fiber.Ctx) error {
	// Take refresh token
	// if token valid
	//   Return tokens
	// else
	//   error

	return c.SendStatus(fiber.StatusNotImplemented)
}
