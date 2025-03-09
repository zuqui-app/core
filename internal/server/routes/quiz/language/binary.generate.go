package language

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/quiz"
	"zuqui-core/internal/quiz/language"
)

func GenerateBinary(c *fiber.Ctx) error {
	config := new(language.BinaryConfig)

	if err := c.BodyParser(config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "malformed quiz setting")
	}

	jsonStr, err := language.GenerateBinary(
		quiz.QuizGeneratorSystemPrompt,
		language.GetBinaryQuestionGenAISchema(),
		language.GetBinaryQuestionPrompt(
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
