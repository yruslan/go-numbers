package game

import (
	"GoNumbers/src/gonumbers/agents"
	"fmt"
)

func play(
	from int,
	to int,
	ask agents.AskAgent,
	guess agents.GuessAgent) (int, error) {

	if err := ask.ThinkOfANewNumber(from, to); err != nil {
		return 0, err
	}

	if err := guess.StartGame(from, to); err != nil {
		return 0, err
	}

	var found = false
	for !found {
		num := guess.NextGuess()

		guessFeedback, err := ask.CheckGuess(num)
		if err != nil {
			return 0, err
		}

		if guessFeedback == 0 {
			fmt.Printf("Congratulations! You have found the number: %d\n", num)
			found = true
		} else {
			if err := guess.GuessFeedback(guessFeedback); err != nil {
				return 0, err
			}
		}
	}

	return 0, nil
}
