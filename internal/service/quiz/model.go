package quiz

type Explanation struct {
	Text    string `json:"text"`
	Example string `json:"example"`
}

type BinaryQuestion struct {
	Question      string      `json:"question"`
	Answer        bool        `json:"answer"`
	Specification string      `json:"specification"`
	Difficulty    string      `json:"difficulty"`
	Explanation   Explanation `json:"explanantion"`
}

type SingleChoiceQuestion struct {
	Question      string      `json:"question"`
	Answer        string      `json:"answer"`
	Incorrects    []string    `json:"incorrects"`
	Specification string      `json:"specification"`
	Difficulty    string      `json:"difficulty"`
	Explanation   Explanation `json:"explanantion"`
}

type MultipleChoiceQuestion struct {
	Question string
	Answers  []int
	Options  []string
}

type ClozeQuestion struct {
	Question []string
	Answers  []string
}

type MatchingQuestion struct {
	A []string
	B []string
}

type OrderingQuestion struct {
	Question string
	Answer   []string
}

type OpenEndedQuestion struct {
	Question string
}
