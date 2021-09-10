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
