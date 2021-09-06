package agents

import (
	"errors"
	"fmt"
)

type GuessAgentPc struct {
	from      int
	to        int
	stateFrom int
	stateTo   int
	lastGuess int
}

func (p *GuessAgentPc) StartGame(from int, to int) error {
	if from >= to {
		return errors.New("StartGame: To should be greater than From")
	}

	p.from = from
	p.to = to
	p.stateFrom = from
	p.stateTo = to
	p.lastGuess = from + (to-from+1)/2

	return nil
}

func (p *GuessAgentPc) NextGuess() int {
	p.lastGuess = p.stateFrom + (p.stateTo-p.stateFrom+1)/2

	fmt.Printf("My next guess: %d\n", p.lastGuess)

	return p.lastGuess
}

func (p *GuessAgentPc) GuessFeedback(guessFeedback int) error {
	if guessFeedback == -1 {
		p.stateFrom = p.lastGuess + 1
	} else if guessFeedback == 1 {
		p.stateTo = p.lastGuess - 1
	} else {
		return fmt.Errorf("Unexpected feedback: %d", guessFeedback)
	}
	return nil
}
