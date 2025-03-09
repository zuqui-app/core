package me

import "github.com/gofiber/fiber/v2"

func Usage(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
