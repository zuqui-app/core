package server

import "github.com/gofiber/fiber/v2"

// Onboard new users
func (s *FiberServer) AuthOnboard(c *fiber.Ctx) error {
	// Take otp and onboard data
	// Finish signup
	// Return tokens

	return c.SendStatus(fiber.StatusNotImplemented)
}
