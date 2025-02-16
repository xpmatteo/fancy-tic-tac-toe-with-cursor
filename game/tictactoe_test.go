package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	game := NewGame()
	assert.Equal(t, "hello", game.Text(), "should return hello text")
}

func TestInitialBoard(t *testing.T) {
	game := NewGame()
	expected := [9]string{"", "", "", "", "", "", "", "", ""}

	assert.Equal(t, expected, game.Board(), "should start with empty board")
}

func TestMakeMove(t *testing.T) {
	tests := []struct {
		name          string
		position      int
		setupMoves    []int
		expectedErr   string
		expectedBoard [9]string
	}{
		{
			name:          "valid move",
			position:      4,
			expectedErr:   "",
			expectedBoard: [9]string{"", "", "", "", "X", "", "", "", ""},
		},
		{
			name:          "cell already occupied",
			position:      4,
			setupMoves:    []int{4},
			expectedErr:   "cell already occupied",
			expectedBoard: [9]string{"", "", "", "", "X", "", "", "", ""},
		},
		{
			name:          "invalid position too high",
			position:      9,
			expectedErr:   "invalid position",
			expectedBoard: [9]string{"", "", "", "", "", "", "", "", ""},
		},
		{
			name:          "invalid position negative",
			position:      -1,
			expectedErr:   "invalid position",
			expectedBoard: [9]string{"", "", "", "", "", "", "", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()

			// Perform any setup moves
			for _, move := range tt.setupMoves {
				_ = game.MakeMove(move)
			}

			// Perform the test move
			err := game.MakeMove(tt.position)

			// Check error
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			// Check board state
			assert.Equal(t, tt.expectedBoard, game.Board())
		})
	}
}
