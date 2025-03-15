package server

import "github.com/gofiber/fiber/v2"

func (s *App) MeProfile(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
