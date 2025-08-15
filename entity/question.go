package entity

type Question struct {
	ID             uint
	Text           string
	PossibleAnswer []PossibleAnswer
	CorrectAnswer  []uint
	Difficulty     QuestionDifficulty
	CategoryID     uint
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	if p >= PossibleNaswerAnswerA && p <= PossibleNaswerAnswerD {
		return true
	}

	return false
}

const (
	PossibleNaswerAnswerA PossibleAnswerChoice = iota + 1
	PossibleNaswerAnswerB
	PossibleNaswerAnswerC
	PossibleNaswerAnswerD
)

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	if q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard {
		return true
	}
	
	return false

}
