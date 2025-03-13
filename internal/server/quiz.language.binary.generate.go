package server

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/quiz/language"
)

func (s *FiberServer) LanguageGenerateBinary(c *fiber.Ctx) error {
	config := new(language.BinaryConfig)

	if err := c.BodyParser(config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "malformed quiz setting")
	}

	jsonStr, err := s.quiz.Language.GenerateBinary(
		s.quiz.Prompt.SystemPrompt(),
		s.quiz.Language.GetBinaryQuestionGenAISchema(),
		s.quiz.Language.GetBinaryQuestionPrompt(
			language.BinaryConfig{
				Language:      "German",
				Category:      "Grammar",
				Specification: "B2",
				Difficulty:    "Hard",
				Amount:        5,
			},
		),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	log.Println(jsonStr)
	return nil
}
