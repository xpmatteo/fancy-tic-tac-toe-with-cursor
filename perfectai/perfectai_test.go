package perfectai

import (
	"testing"
	"tictactoe/game"

	"github.com/stretchr/testify/assert"
)

func TestPerfectAI_ChooseMove(t *testing.T) {
	tests := []struct {
		name     string
		board    [9]string
		expected int
	}{
		{
			name: "block X's winning move",
			board: [9]string{
				"X", "X", "",
				"", "", "",
				"", "", "",
			},
			expected: 2, // must block at position 2
		},
		{
			name: "force draw when possible",
			board: [9]string{
				"X", "X", "O",
				"O", "O", "",
				"X", "", "X",
			},
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := game.FromBoard(test.board)
			ai := NewPerfectAI()
			move := ai.ChooseMove(g)
			assert.Equal(t, test.expected, move)
		})
	}
}
