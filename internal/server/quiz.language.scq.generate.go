package server

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/quiz/language"
)

func (s *FiberServer) LanguageGenerateSCQ(c *fiber.Ctx) error {
	config := new(language.SCQConfig)

	if err := c.BodyParser(config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "malformed quiz setting")
	}

	jsonStr, err := s.quiz.Language.GenerateSCQ(
		s.quiz.Prompt.SystemPrompt(),
		s.quiz.Language.GetSingleChoiceQuestionGenAISchema(),
		s.quiz.Language.GetSingleChoiceQuestionPrompt(
			language.SCQConfig{
				Language:      "German",
				Category:      "Grammar",
				Specification: "B2",
				Difficulty:    "Medium",
				ChoiceAmount:  4,
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
