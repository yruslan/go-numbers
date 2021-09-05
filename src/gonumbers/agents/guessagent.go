package agents

type GuessAgent interface {
	StartGame(int, int) error

	NextGuess() int

	GuessFeedback(guessFeedback int) error
}
