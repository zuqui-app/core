package language

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"

	"zuqui-core/internal/quiz"
)

const BinaryQuestionThoughtProcess = "### Thought Process for Binary Question ###\n" +
	"1. Identify the key knowledge point to be tested.\n" +
	"2. Select a specific aspect of the topic to focus on.\n" +
	"3. Craft a True/False question that targets that aspect, ensuring it has a clear, singular correct answer.\n"

type BinaryConfig struct {
	Language      string `json:"language"      xml:"language"      form:"language"`
	Category      string `json:"category"      xml:"category"      form:"category"`
	Specification string `json:"specification" xml:"specification" form:"specification"`
	Difficulty    string `json:"difficulty"    xml:"difficulty"    form:"difficulty"`
	Amount        int    `json:"amount"        xml:"amount"        form:"amount"`
}

func GetBinaryQuestionPrompt(config BinaryConfig) string {
	return fmt.Sprintf(
		"Generate True-or-False questions that test knowledge of a specific language.\n"+
			quiz.DesirableQuizCriteria+
			BinaryQuestionThoughtProcess+
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
}

func GetBinaryQuestionGenAISchema() *genai.Schema {
	return &genai.Schema{
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
}

const SingleChoiceQuestionThoughtProcess = "### Thought Process of Individual Question ###\n" +
	"1. Identify a knowledge point related to the quiz configuration.\n" +
	"2. Choose an aspect to focus the question on.\n" +
	"3. Formulate a question that tests the knowledge within that aspect, ensuring only one correct and definite answer.\n" +
	"4. Generate plausible yet definitely incorrect answers, based on common misconceptions or slight variations of the correct answer."

type SCQConfig struct {
	Language      string `json:"language"      xml:"language"      form:"language"`
	Category      string `json:"category"      xml:"category"      form:"category"`
	Specification string `json:"specification" xml:"specification" form:"specification"`
	Difficulty    string `json:"difficulty"    xml:"difficulty"    form:"difficulty"`
	ChoiceAmount  int    `json:"choiceAmount"  xml:"choiceAmount"  form:"choiceAmount"`
	Amount        int    `json:"amount"        xml:"amount"        form:"amount"`
}

func GetSingleChoiceQuestionPrompt(config SCQConfig) string {
	return fmt.Sprintf(
		"Generate Single-Choice questions that tests knowledge of a specific language\n"+
			quiz.DesirableQuizCriteria+
			BinaryQuestionThoughtProcess+
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
}

func GetSingleChoiceQuestionGenAISchema() *genai.Schema {
	return &genai.Schema{
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
}
