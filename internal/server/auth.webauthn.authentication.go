package server

import "github.com/gofiber/fiber/v2"

// Authenticate with Passkey
func (s *App) WebAuthnAuthentication(c *fiber.Ctx) error {
	// challenge
	// if challenge success
	//   find user with corresponding public key
	//   return tokens
	// else
	//   error

	return c.SendStatus(fiber.StatusNotImplemented)
}
