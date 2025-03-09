package auth

import "github.com/gofiber/fiber/v2"

// Authenticate with Passkey, requires existing account
func WebAuthnRegistration(c *fiber.Ctx) error {
	// take access token and challenge
	// if access valid and challenge succeeded
	//   bind public key to user account
	// else
	//   error

	return c.SendStatus(fiber.StatusNotImplemented)
}
