package agents

import (
	"fmt"
	"math/rand"
	"time"
)

type AskAgentPc struct {
	number int
}

func (p *AskAgentPc) ThinkOfANewNumber(from int, to int) error {
	if from >= to {
		return fmt.Errorf("ThinkOfANewNumber(): To should be greater than From")
	}

	rand.Seed(time.Now().UnixNano())

	p.number = rand.Intn(to-from) + from

	fmt.Printf("I have guessed a number between %d and %d. Try to guess it.\n", from, to)
	return nil
}

func (p *AskAgentPc) CheckGuess(guess int) int {
	if guess < p.number {
		fmt.Printf("Too small.\n")
		return -1
	} else if guess > p.number {
		fmt.Printf("Too big.\n")
		return 1
	} else {
		fmt.Printf("Yes! That's the number.\n")
		return 0
	}
}
