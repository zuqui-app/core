package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"

	"zuqui-core/internal"
)

type language struct {
	client *genai.Client
}

type LanguageBinaryConfig struct {
	Language      string `json:"language"      xml:"language"      form:"language"      validate:"required"`
	Category      string `json:"category"      xml:"category"      form:"category"      validate:"required_with=Specification"`
	Specification string `json:"specification" xml:"specification" form:"specification"`
	Difficulty    string `json:"difficulty"    xml:"difficulty"    form:"difficulty"    validate:"required"`
	Amount        int    `json:"amount"        xml:"amount"        form:"amount"        validate:"required"`
}

func (q *language) Binary(config LanguageBinaryConfig) (*[]BinaryQuestion, error) {
	if err := internal.Validate.Struct(config); err != nil {
		return nil, err
	}

	model := q.client.GenerativeModel("gemini-2.0-flash-lite")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	model.ResponseSchema = binaryQuestionsSchema
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	prompt := fmt.Sprintf(
		"Generate True-or-False questions that test knowledge of a specific language.\n"+
			quizStandard+
			binaryQuestionThoughtProcess+
			"### Quiz Configurations ###\n"+
			"Language: %s\n"+
			"Category: %s\n"+
			"Further specification: %s\n"+
			"General difficulty for the questions: %s\n"+
			"Amount of questions to generate: %d",
		config.Language,
		config.Category,
		config.Specification,
		config.Difficulty,
		config.Amount,
	)

	ctx := context.Background()
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	var questions []BinaryQuestion
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var part_of_questions []BinaryQuestion
			if err := json.Unmarshal([]byte(txt), &part_of_questions); err != nil {
				log.Println(err)
			}
			questions = append(questions, part_of_questions...)
		}
	}

	return &questions, nil
}

type LanguageSingleChoiceConfig struct {
	Language      string `json:"language"      xml:"language"      form:"language"      validate:"required"`
	Category      string `json:"category"      xml:"category"      form:"category"`
	Specification string `json:"specification" xml:"specification" form:"specification"`
	Difficulty    string `json:"difficulty"    xml:"difficulty"    form:"difficulty"    validate:"required"`
	ChoiceAmount  int    `json:"choiceAmount"  xml:"choiceAmount"  form:"choiceAmount"  validate:"required"`
	Amount        int    `json:"amount"        xml:"amount"        form:"amount"        validate:"required"`
}

func (q *language) SingleChoice(
	config LanguageSingleChoiceConfig,
) (*[]SingleChoiceQuestion, error) {
	if err := internal.Validate.Struct(config); err != nil {
		return nil, err
	}

	model := q.client.GenerativeModel("gemini-2.0-flash-lite")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	model.ResponseSchema = singleChoiceQuestionsSchema
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	prompt := fmt.Sprintf(
		"Generate Single-Choice questions that tests knowledge of a specific language\n"+
			quizStandard+
			binaryQuestionThoughtProcess+
			"### Quiz Configurations ###\n"+
			"Language: %s\n"+
			"Category: %s\n"+
			"Further specification: %s\n"+
			"General difficulty for the questions: %s\n"+
			"Amount of choices for each question: %d\n"+
			"Amount of questions to generate: %d",
		config.Language,
		config.Category,
		config.Specification,
		config.Difficulty,
		config.ChoiceAmount,
		config.Amount,
	)

	ctx := context.Background()
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	var questions []SingleChoiceQuestion
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var part_of_questions []SingleChoiceQuestion
			if err := json.Unmarshal([]byte(txt), &part_of_questions); err != nil {
				log.Println(err)
			}
			questions = append(questions, part_of_questions...)
		}
	}

	return &questions, nil
}
