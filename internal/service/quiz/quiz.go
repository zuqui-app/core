package quiz

import "github.com/google/generative-ai-go/genai"

type Service struct {
	Language language
	Math     math
}

func New(client *genai.Client) *Service {
	return &Service{
		Language: language{client},
		Math:     math{client},
	}
}
