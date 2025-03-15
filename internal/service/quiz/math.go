package quiz

import "github.com/google/generative-ai-go/genai"

type math struct {
	client *genai.Client
}

func (q *math) Binary() {}

func (q *math) SCQ() {}
