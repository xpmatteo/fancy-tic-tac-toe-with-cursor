package game

import (
	"fmt"
)

type Game struct {
	board   [9]string
	xIsNext bool
}

func NewGame() *Game {
	return &Game{
		board:   [9]string{"", "", "", "", "", "", "", "", ""},
		xIsNext: true, // X plays first
	}
}

func (g *Game) Board() [9]string {
	return g.board
}

func (g *Game) MakeMove(position int) error {
	if position < 0 || position >= len(g.board) {
		return fmt.Errorf("invalid position")
	}
	if g.board[position] != "" {
		return fmt.Errorf("cell already occupied")
	}
	if g.xIsNext {
		g.board[position] = "X"
	} else {
		g.board[position] = "O"
	}
	g.xIsNext = !g.xIsNext // Switch turns
	return nil
}

func (g *Game) Text() string {
	if winner := g.Winner(); winner != "" {
		return winner + " has won!"
	}
	if g.xIsNext {
		return "X to move"
	}
	return "O to move"
}

func (g *Game) Winner() string {
	if g.hasWinningLine() {
		return "X" // X wins if there's any winning combination
	}

	if g.IsDraw() {
		return "O" // O wins if the game is a draw
	}

	return "" // No winner
}

func (g *Game) IsDraw() bool {
	// Count empty cells
	emptyCount := 0
	emptyPos := -1
	for pos, cell := range g.board {
		if cell == "" {
			emptyCount++
			emptyPos = pos
		}
	}

	// If more than one empty cell, not a draw
	if emptyCount > 1 {
		return false
	}

	// If board is full and no winner, it's a draw
	if emptyCount == 0 {
		return !g.hasWinningLine()
	}

	// If exactly one empty cell, check if filling it with X leads to a win
	g.board[emptyPos] = "X"
	hasWin := g.hasWinningLine()
	g.board[emptyPos] = "" // Reset
	return !hasWin
}

func (g *Game) AvailableMoves() []int {
	return nil
}

func (g *Game) findWinningLine() ([3]int, bool) {
	winningCombinations := [][3]int{
		{0, 1, 2}, // Top row
		{3, 4, 5}, // Middle row
		{6, 7, 8}, // Bottom row
		{0, 3, 6}, // Left column
		{1, 4, 7}, // Middle column
		{2, 5, 8}, // Right column
		{0, 4, 8}, // Diagonal \
		{2, 4, 6}, // Diagonal /
	}
	for _, combo := range winningCombinations {
		if g.isWinningLine(combo) {
			return combo, true
		}
	}
	return [3]int{}, false
}

func (g *Game) hasWinningLine() bool {
	_, ok := g.findWinningLine()
	return ok
}

func (g *Game) isWinningLine(combo [3]int) bool {
	return g.board[combo[0]] != "" && g.board[combo[0]] == g.board[combo[1]] && g.board[combo[1]] == g.board[combo[2]]
}
