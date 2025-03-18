package quiz

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"zuqui/internal"
	"zuqui/internal/service/quiz"
)

func LanguageGenerateBinary(qs *quiz.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var config quiz.LanguageBinaryConfig

		if err := c.BodyParser(&config); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "malformed request")
		}

		questions, err := qs.Language.Binary(config)
		if err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				return c.Status(fiber.StatusBadRequest).JSON(internal.ValidationErrorsToMap(errors))
			}

			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.JSON(*questions)
	}
}
