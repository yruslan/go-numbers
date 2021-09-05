package agents

type AskAgent interface {
	ThinkOfANewNumber(int, int) error

	CheckGuess(guess int) (int, error)
}
