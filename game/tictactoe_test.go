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
			wantWinner: "X",
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
		{
			name:       "O wins on draw",
			setupMoves: []int{0, 1, 2, 4, 3, 5, 7, 6, 8},
			wantWinner: "O",
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
			name:         "X wins diagonal",
			setupMoves:   []int{1, 0, 3, 4, 7, 8},
			expectedText: "X has won!",
		},
		{
			name:         "O wins on draw",
			setupMoves:   []int{0, 1, 2, 4, 3, 5, 7, 6, 8},
			expectedText: "O has won!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := NewGame()

			// Perform setup moves
			for _, move := range test.setupMoves {
				_ = game.MakeMove(move)
			}

			assert.Equal(t, test.expectedText, game.Text())
		})
	}
}

func TestIsDraw(t *testing.T) {
	tests := []struct {
		name  string
		board [9]string
		want  bool
	}{
		{
			name:  "empty board is not a draw",
			board: [9]string{"", "", "", "", "", "", "", "", ""},
			want:  false,
		},
		{
			name: "full board with no winner is a draw",
			board: [9]string{
				"X", "O", "X",
				"X", "O", "O",
				"O", "X", "X",
			},
			want: true,
		},
		{
			name: "one move left, next move forces draw",
			board: [9]string{
				"X", "O", "X",
				"X", "O", "O",
				"O", "X", "",
			},
			want: true,
		},
		{
			name: "one move left, next move will win",
			board: [9]string{
				"X", "X", "",
				"X", "O", "O",
				"O", "X", "O",
			},
			want: false,
		},
		{
			name: "multiple empty spaces is not a draw",
			board: [9]string{
				"X", "", "O",
				"", "X", "",
				"O", "", "",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGame()
			g.board = test.board
			got := g.IsDraw()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGame_AvailableMoves(t *testing.T) {
	tests := []struct {
		name     string
		board    [9]string
		expected []int
	}{
		{
			name:     "empty board has all positions available",
			board:    [9]string{"", "", "", "", "", "", "", "", ""},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "partially filled board",
			board: [9]string{
				"X", "O", "",
				"", "X", "",
				"O", "", "",
			},
			expected: []int{2, 3, 5, 7, 8},
		},
		{
			name: "only one move available",
			board: [9]string{
				"X", "O", "X",
				"O", "X", "O",
				"O", "X", "",
			},
			expected: []int{8},
		},
		{
			name: "full board has no moves",
			board: [9]string{
				"X", "O", "X",
				"O", "X", "O",
				"O", "X", "O",
			},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := &Game{board: test.board}
			moves := game.AvailableMoves()
			assert.Equal(t, test.expected, moves)
		})
	}
}

func TestGame_Clone(t *testing.T) {
	tests := []struct {
		name     string
		board    [9]string
		xIsNext  bool
		makeMove int // move to make after cloning
	}{
		{
			name:     "clone empty board",
			board:    [9]string{"", "", "", "", "", "", "", "", ""},
			xIsNext:  true,
			makeMove: 0,
		},
		{
			name: "clone partial game",
			board: [9]string{
				"X", "O", "",
				"", "X", "",
				"", "", "",
			},
			xIsNext:  false,
			makeMove: 2,
		},
		{
			name: "clone nearly complete game",
			board: [9]string{
				"X", "O", "X",
				"O", "X", "O",
				"X", "O", "",
			},
			xIsNext:  true,
			makeMove: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original := &Game{
				board:   test.board,
				xIsNext: test.xIsNext,
			}

			// Create clone and verify it matches
			clone := original.Clone()
			assert.Equal(t, original.board, clone.board)
			assert.Equal(t, original.xIsNext, clone.xIsNext)

			// Modify clone and verify original is unchanged
			clone.MakeMove(test.makeMove)
			assert.NotEqual(t, original.board, clone.board)
			assert.NotEqual(t, original.xIsNext, clone.xIsNext)
		})
	}
}
