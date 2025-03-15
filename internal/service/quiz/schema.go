package quiz

import "github.com/google/generative-ai-go/genai"

var binaryQuestionsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type:     genai.TypeObject,
		Required: []string{"question", "answer", "specification", "difficulty", "explanation"},
		Properties: map[string]*genai.Schema{
			"question": {
				Type: genai.TypeString,
			},
			"answer": {
				Type: genai.TypeBoolean,
			},
			"specification": {
				Type:        genai.TypeString,
				Description: "What specific aspect the question is about",
			},
			"difficulty": {
				Type: genai.TypeString,
				Enum: []string{"Easy", "Medium", "Hard"},
			},
			"explanation": {
				Type:        genai.TypeObject,
				Description: "A brief explanation on the correct answer",
				Required:    []string{"text", "example"},
				Properties: map[string]*genai.Schema{
					"text": {
						Type:        genai.TypeString,
						Description: "The explanation",
					},
					"example": {
						Type:        genai.TypeString,
						Description: "A brief example to support the explanation",
					},
				},
			},
		},
	},
}

var singleChoiceQuestionsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type: genai.TypeObject,
		Required: []string{
			"question",
			"answer",
			"incorrects",
			"specification",
			"difficulty",
			"explanation",
		},
		Properties: map[string]*genai.Schema{
			"question": {
				Type: genai.TypeString,
			},
			"answer": {
				Type: genai.TypeString,
			},
			"incorrects": {
				Type:        genai.TypeArray,
				Description: "Incorrect responses to the question",
				Items: &genai.Schema{
					Type: genai.TypeString,
				},
			},
			"specification": {
				Type:        genai.TypeString,
				Description: "What specific aspect the question is about",
			},
			"difficulty": {
				Type: genai.TypeString,
			},
			"explanation": {
				Type:        genai.TypeObject,
				Description: "A brief explanation on the correct answer",
				Required:    []string{"text", "example"},
				Properties: map[string]*genai.Schema{
					"text": {
						Type:        genai.TypeString,
						Description: "The explanation",
					},
					"example": {
						Type:        genai.TypeString,
						Description: "A brief example to support the explanation",
					},
				},
			},
		},
	},
}
