package quiz

import (
	"github.com/google/generative-ai-go/genai"

	"zuqui-core/internal/quiz/language"
	"zuqui-core/internal/quiz/prompt"
)

type Generator struct {
	Language *language.Generator
	Prompt   *prompt.Prompt
}

func New(ai *genai.Client) *Generator {
	prompt := &prompt.Prompt{}
	return &Generator{
		Language: language.New(ai, prompt),
		Prompt:   prompt,
	}
}
