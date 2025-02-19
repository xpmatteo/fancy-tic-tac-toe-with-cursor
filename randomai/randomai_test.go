package randomai

import (
	"testing"
	"tictactoe/game"

	"github.com/stretchr/testify/assert"
)

// StubRNG implements RandomGenerator for testing
type StubRNG struct {
	nextValue int
}

func (s *StubRNG) NextValue(n int) int {
	return s.nextValue
}

func TestRandomAI_ChooseMove(t *testing.T) {
	tests := []struct {
		name       string
		board      [9]string
		stubReturn int
		expected   int
	}{
		{
			name:       "empty board - should use RNG value directly",
			board:      [9]string{"", "", "", "", "", "", "", "", ""},
			stubReturn: 4,
			expected:   4,
		},
		{
			name: "partially filled board - should use RNG value to index available moves",
			board: [9]string{
				"X", "", "",
				"", "O", "",
				"", "", "X",
			},
			stubReturn: 2,
			expected:   3, // third available move (1,2,3,5,6,7)
		},
		{
			name: "only one move available - should return the only move regardless of RNG",
			board: [9]string{
				"X", "O", "X",
				"O", "X", "O",
				"O", "X", "",
			},
			stubReturn: 0,
			expected:   8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := game.FromBoard(test.board)
			rng := &StubRNG{nextValue: test.stubReturn}
			ai := NewRandomAI(rng)
			move := ai.ChooseMove(g)

			assert.Equal(t, test.expected, move)
		})
	}
}
