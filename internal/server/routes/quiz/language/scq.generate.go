package language

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/quiz"
	"zuqui-core/internal/quiz/language"
)

func GenerateSCQ(c *fiber.Ctx) error {
	config := new(language.SCQConfig)

	if err := c.BodyParser(config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "malformed quiz setting")
	}

	jsonStr, err := language.GenerateSCQ(
		quiz.QuizGeneratorSystemPrompt,
		language.GetSingleChoiceQuestionGenAISchema(),
		language.GetSingleChoiceQuestionPrompt(
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
