package user

import "github.com/gofiber/fiber/v2"

func Usage() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
