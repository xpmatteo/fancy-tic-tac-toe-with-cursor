package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			name:          "X plays first",
			position:      4,
			expectedErr:   "",
			expectedBoard: [9]string{"", "", "", "", "X", "", "", "", ""},
		},
		{
			name:          "O plays second",
			position:      0,
			setupMoves:    []int{4},
			expectedErr:   "",
			expectedBoard: [9]string{"O", "", "", "", "X", "", "", "", ""},
		},
		{
			name:          "X plays third",
			position:      8,
			setupMoves:    []int{4, 0},
			expectedErr:   "",
			expectedBoard: [9]string{"O", "", "", "", "X", "", "", "", "X"},
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

func TestWinner(t *testing.T) {
	tests := []struct {
		name       string
		setupMoves []int
		wantWinner string
	}{
		{
			name:       "no winner on empty board",
			setupMoves: []int{},
			wantWinner: "",
		},
		{
			name:       "X wins top row",
			setupMoves: []int{0, 3, 1, 4, 2},
			wantWinner: "X",
		},
		{
			name:       "O wins diagonal",
			setupMoves: []int{1, 0, 3, 4, 7, 8},
			wantWinner: "O",
		},
		{
			name:       "no winner in ongoing game",
			setupMoves: []int{0, 1, 2, 3},
			wantWinner: "",
		},
		{
			name:       "X wins vertical",
			setupMoves: []int{1, 0, 4, 3, 7},
			wantWinner: "X",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := NewGame()
			for _, move := range test.setupMoves {
				_ = game.MakeMove(move)
			}
			assert.Equal(t, test.wantWinner, game.Winner())
		})
	}
}

func TestText(t *testing.T) {
	tests := []struct {
		name         string
		setupMoves   []int
		expectedText string
	}{
		{
			name:         "X's turn at start",
			setupMoves:   []int{},
			expectedText: "X to move",
		},
		{
			name:         "O's turn after X plays",
			setupMoves:   []int{4},
			expectedText: "O to move",
		},
		{
			name:         "X's turn after O plays",
			setupMoves:   []int{4, 0},
			expectedText: "X to move",
		},
		{
			name:         "X wins top row",
			setupMoves:   []int{0, 3, 1, 4, 2},
			expectedText: "X has won!",
		},
		{
			name:         "O wins diagonal",
			setupMoves:   []int{1, 0, 3, 4, 7, 8},
			expectedText: "O has won!",
		},
		{
			name:         "game is a draw",
			setupMoves:   []int{0, 1, 2, 4, 3, 5, 7, 6, 8},
			expectedText: "Game is a draw!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := NewGame()

			for _, move := range test.setupMoves {
				_ = game.MakeMove(move)
			}

			assert.Equal(t, test.expectedText, game.Text())
		})
	}
}

func TestIsDraw(t *testing.T) {
	tests := []struct {
		name       string
		setupMoves []int
		wantDraw   bool
	}{
		{
			name:       "empty board is not a draw",
			setupMoves: []int{},
			wantDraw:   false,
		},
		{
			name:       "game in progress is not a draw",
			setupMoves: []int{0, 1, 2, 3},
			wantDraw:   false,
		},
		{
			name:       "won game is not a draw",
			setupMoves: []int{0, 3, 1, 4, 2},
			wantDraw:   false,
		},
		{
			name:       "draw when board is full with no winner",
			setupMoves: []int{0, 1, 2, 4, 3, 5, 7, 6, 8},
			wantDraw:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := NewGame()
			for _, move := range test.setupMoves {
				_ = game.MakeMove(move)
			}
			assert.Equal(t, test.wantDraw, game.IsDraw())
		})
	}
}
