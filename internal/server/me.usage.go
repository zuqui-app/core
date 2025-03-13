package server

import "github.com/gofiber/fiber/v2"

func (s *FiberServer) MeUsage(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
