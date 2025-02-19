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
			name: "take center when X starts in corner",
			board: [9]string{
				"X", "", "",
				"", "", "",
				"", "", "",
			},
			expected: 4, // center is best response
		},
		{
			name: "force draw when possible",
			board: [9]string{
				"X", "", "",
				"", "O", "",
				"", "", "X",
			},
			expected: 1, // taking corner would allow X to win
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
