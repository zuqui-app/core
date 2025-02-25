package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type QuizSetting struct {
	Subject    string `json:"subject"    xml:"subject"    form:"subject"`
	Category   string `json:"category"   xml:"category"   form:"category"`
	Range      string `json:"range"      xml:"range"      form:"range"`
	Difficulty string `json:"difficulty" xml:"difficulty" form:"difficulty"`
	QuizAmount int    `json:"quizAmount" xml:"quizAmount" form:"quizAmount"`
}

func (s *FiberServer) CreateQuizHandler(c *fiber.Ctx) error {
	quizSetting := new(QuizSetting)

	if err := c.BodyParser(quizSetting); err != nil {
		return err
	}

	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(
				"You are a dedicated and powerful quiz generator. Accuracy is your top priority, and you take great care in ensuring that every quiz you create is factually correct and meaningful. Your goal is not just to generate quizzes, but to educate quiz takers by providing varied and comprehensive questions. You are committed to delivering quizzes that cover a wide range of knowledge and concepts, ensuring that each quiz serves to enhance the understanding of the subject matter.",
			),
		},
	}

	model.ResponseSchema = &genai.Schema{
		Type:     genai.TypeObject,
		Required: []string{"response"},

		Properties: map[string]*genai.Schema{
			"response": {
				Type: genai.TypeArray,
				Items: &genai.Schema{
					Type:     genai.TypeObject,
					Required: []string{"answer", "false_answers", "question", "category"},
					Properties: map[string]*genai.Schema{
						"answer": {
							Type: genai.TypeString,
						},
						"false_answers": {
							Type: genai.TypeArray,
							Items: &genai.Schema{
								Type: genai.TypeString,
							},
						},
						"question": {
							Type: genai.TypeString,
						},
						"category": {
							Type: genai.TypeString,
						},
						"subcategory": {
							Type: genai.TypeString,
						},
					},
				},
			},
		},
	}

	session := model.StartChat()
	prompt := fmt.Sprintf(
		"Generate quizzes based on the information and criteria below.\n\n"+
			"### Criteria for Desirable Quizzes ###\n"+
			"1. The question should focus on testing the user's understanding of the concept or grammar rule in context.\n"+
			"2. The quizzes should cover a wide range of concepts within the subject.\n"+
			"3. The quizzes should span across different aspects of the subject, ensuring diversity in question types and content.\n"+
			"4. The quizzes should come from various angles, avoiding repetitive sentence structures or question types.\n"+
			"5. The quizzes should not hint at the answer directly in the question, and should not make the correct answer obvious by mentioning a key element of it.\n"+
			"\n"+
			"### Thought Process ###\n"+
			"Let's break it down step by step:\n"+
			"1. Identify a fact or knowledge point related to the information and criteria.\n"+
			"2. Double-check that the fact or knowledge is accurate and error-free.\n"+
			"3. Choose a category or aspect to focus the quiz on.\n"+
			"4. Formulate a question that tests the knowledge within that category, ensuring only one correct and definite answer.\n"+
			"5. Verify that the question is related to the correct answer.\n"+
			"6. Ensure that the question appropriately tests the user's understanding of the intended aspect.\n"+
			"7. Double-check that the question results in a clear, correct answer.\n"+
			"8. Generate 3 plausible yet incorrect answers, based on common misconceptions or slight variations of the correct answer.\n"+
			"9. Confirm that the false answers are plausible but definitely incorrect.\n"+
			"\n"+
			"### Information and Criteria ###\n"+
			"Quiz Subject: %s\n"+
			"Quiz Category: %s\n"+
			"Quiz Range: %s\n"+
			"Quiz Difficulty: %s\n"+
			"Amount of quizzes generated: %d",
		quizSetting.Subject,
		quizSetting.Category,
		quizSetting.Range,
		quizSetting.Difficulty,
		quizSetting.QuizAmount,
	)

	resp, err := session.SendMessage(
		ctx,
		genai.Text(prompt),
	)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}

	return c.SendString("Getcha")
}
