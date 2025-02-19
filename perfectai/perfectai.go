package perfectai

import "tictactoe/game"

type PerfectAI struct{}

func NewPerfectAI() *PerfectAI {
	return &PerfectAI{}
}

func (ai *PerfectAI) ChooseMove(g *game.Game) int {
	bestScore := -1000
	bestMove := -1

	for _, move := range g.AvailableMoves() {
		// Try this move on a clone
		clone := g.Clone()
		clone.MakeMove(move)

		// Evaluate position after X's best response
		score := -ai.minimax(clone, false)

		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func (ai *PerfectAI) minimax(g *game.Game, isMax bool) int {
	winner := g.Winner()
	if winner == "X" {
		return -100 // O loses (X achieved a winning position)
	}
	if winner == "O" {
		return 100 // O wins (forced a draw)
	}

	moves := g.AvailableMoves()
	if len(moves) == 0 {
		panic("No moves available") // Should never happen as Winner() would have returned something
	}

	if isMax {
		// X's turn - X tries to maximize score
		bestScore := -1000
		for _, move := range moves {
			clone := g.Clone()
			clone.MakeMove(move)
			score := ai.minimax(clone, false)
			if score > bestScore {
				bestScore = score
			}
		}
		return bestScore
	} else {
		// O's turn - O tries to minimize score
		bestScore := 1000
		for _, move := range moves {
			clone := g.Clone()
			clone.MakeMove(move)
			score := ai.minimax(clone, true)
			if score < bestScore {
				bestScore = score
			}
		}
		return bestScore
	}
}
