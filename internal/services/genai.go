package services

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"zuqui-core/internal"
)

var GenaiClient *genai.Client

func init() {
	client, err := genai.NewClient(
		context.Background(),
		option.WithAPIKey(internal.Env.GEMINI_API_KEY),
	)
	if err != nil {
		log.Fatalf("Error creating genai client: %v", err)
	}

	GenaiClient = client
}
