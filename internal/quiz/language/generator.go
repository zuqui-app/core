package language

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"

	"zuqui-core/internal/quiz/prompt"
)

type Generator struct {
	ai     *genai.Client
	prompt *prompt.Prompt
}

func New(ai *genai.Client, prompt *prompt.Prompt) *Generator {
	return &Generator{ai, prompt}
}

func (g *Generator) GenerateBinary(
	systemPrompt string,
	schema *genai.Schema,
	prompt string,
) (string, error) {
	model := g.ai.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	model.ResponseSchema = schema
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	ctx := context.Background()
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println("Error generating content")
		return "", errors.New("Error generating content")
	}

	jsonStr := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	return jsonStr, nil
}

func (g *Generator) GenerateSCQ(
	systemPrompt string,
	schema *genai.Schema,
	prompt string,
) (string, error) {
	model := g.ai.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	model.ResponseSchema = schema
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	ctx := context.Background()
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println("Error generating content")
		return "", errors.New("Error generating content")
	}

	jsonStr := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	return jsonStr, nil
}
