package main

import (
	"GoNumbers/src/gonumbers/agents"
	"GoNumbers/src/gonumbers/game"
	"fmt"
)

func main() {
	askagent := agents.AskAgentPc{}
	guessagent := agents.GuessAgentPc{}

	if _, err := game.Play(1, 100, &askagent, &guessagent); err != nil {
		fmt.Printf("An error has occurred: %s", err)
	}
}
