package quiz

type BinaryQuestion struct {
	Question string
	Answer   bool
}

type SingleChoiceQuestion struct {
	Question string
	Answer   int
	Options  []string
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
