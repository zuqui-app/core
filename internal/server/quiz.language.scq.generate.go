package server

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal"
	"zuqui-core/internal/service/quiz"
)

func (s *App) LanguageGenerateSCQ(c *fiber.Ctx) error {
	config := new(quiz.LanguageSingleChoiceConfig)

	if err := c.BodyParser(config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "malformed quiz setting")
	}

	questions, err := s.quiz.Language.SingleChoice(*config)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return c.Status(fiber.StatusBadRequest).JSON(internal.ValidationErrorsToMap(errors))
		}

		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(questions)
}
