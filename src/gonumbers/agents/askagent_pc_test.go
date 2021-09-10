package agents

import (
	"testing"
)

func TestThinkOfANumber(t *testing.T) {
	askAgent := AskAgentPc{}

	for i := 1; i < 100; i++ {
		err := askAgent.ThinkOfANewNumber(1, 5)
		if err != nil {
			t.Errorf("Got unexpected error from the agent: %s", err)
		}
		if askAgent.number < 1 || askAgent.number > 5 {
			t.Errorf("Got out of range number %q, expected from 1 to 10", askAgent.number)
		}
	}
}

func TestCheckGuess(t *testing.T) {
	askAgent := AskAgentPc{number: 5}

	guess1 := askAgent.CheckGuess(0)
	if guess1 != -1 {
		t.Errorf("CheckGuess: Got %d, expected -1", guess1)
	}

	guess2 := askAgent.CheckGuess(4)
	if guess2 != -1 {
		t.Errorf("CheckGuess: Got %d, expected -1", guess2)
	}

	guess3 := askAgent.CheckGuess(5)
	if guess3 != 0 {
		t.Errorf("CheckGuess: Got %d, expected 0", guess3)
	}

	guess4 := askAgent.CheckGuess(6)
	if guess4 != 1 {
		t.Errorf("CheckGuess: Got %d, expected 1", guess4)
	}

	guess5 := askAgent.CheckGuess(12)
	if guess5 != 1 {
		t.Errorf("CheckGuess: Got %d, expected 1", guess5)
	}
}
