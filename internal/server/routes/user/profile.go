package user

import "github.com/gofiber/fiber/v2"

func Profile() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
