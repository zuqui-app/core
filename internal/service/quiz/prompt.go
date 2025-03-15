package quiz

const (
	systemPrompt = "You are a dedicated and powerful question generator. Accuracy is your top priority, and you take great care in ensuring that every question you create is factually correct and meaningful. Your goal is generate varied and comprehensive educative questions."
	quizStandard = "### Criteria for Desirable Questions ###\n" +
		"- The quizzes should be brief and direct.\n" +
		"- The quizzes should span across different aspects of the subject, ensuring diversity in question types and content.\n" +
		"- The quizzes should come from various angles, avoiding repetitive sentence structures or question types.\n" +
		"- The quizzes should not hint at the answer directly in the question, and should not make the correct answer obvious by mentioning a key element of it.\n"
)

const binaryQuestionThoughtProcess = "### Thought Process for Binary Question ###\n" +
	"1. Identify the key knowledge point to be tested.\n" +
	"2. Select a specific aspect of the topic to focus on.\n" +
	"3. Craft a True/False question that targets that aspect, ensuring it has a clear, singular correct answer.\n"

const singleChoiceQuestionThoughtProcess = "### Thought Process for Single Choice Question ###\n" +
	"1. Identify the key knowledge point to be tested.\n" +
	"2. Select a specific aspect of the topic to focus on.\n" +
	"3. Craft a single choice question that targets that aspect, ensuring it has a clear, singular correct answer.\n"
