package prompt

const (
	SystemPrompt = "You are a dedicated and powerful question generator. Accuracy is your top priority, and you take great care in ensuring that every question you create is factually correct and meaningful. Your goal is generate varied and comprehensive educative questions."
	QuizStandard = "### Criteria for Desirable Questions ###\n" +
		"- The quizzes should be brief and direct.\n" +
		"- The quizzes should span across different aspects of the subject, ensuring diversity in question types and content.\n" +
		"- The quizzes should come from various angles, avoiding repetitive sentence structures or question types.\n" +
		"- The quizzes should not hint at the answer directly in the question, and should not make the correct answer obvious by mentioning a key element of it.\n"
)

type Prompt struct{}

func (p *Prompt) SystemPrompt() string {
	return SystemPrompt
}

func (p *Prompt) QuizStandard() string {
	return QuizStandard
}
