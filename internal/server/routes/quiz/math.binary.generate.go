package quiz

import "github.com/gofiber/fiber/v2"

func MathGenerateBinary() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
