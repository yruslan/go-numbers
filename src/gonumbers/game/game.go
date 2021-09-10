package game

import (
	"GoNumbers/src/gonumbers/agents"
	"fmt"
)

func Play(
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
	var attempt = 1
	for !found {
		fmt.Printf("-- Attempt #%d --\n", attempt)

		num := guess.NextGuess()
		guessFeedback := ask.CheckGuess(num)

		if guessFeedback == 0 {
			fmt.Printf("Congratulations! You have found the number: %d\n", num)
			found = true
		} else {
			if err := guess.GuessFeedback(guessFeedback); err != nil {
				return 0, err
			}
		}

		attempt += 1
	}

	return 0, nil
}
