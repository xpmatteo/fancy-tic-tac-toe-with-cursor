package randomai

import "tictactoe/game"

type RandomGenerator interface {
	NextValue(n int) int
}

type RandomAI struct {
	generator RandomGenerator
}

func NewRandomAI(generator RandomGenerator) *RandomAI {
	return &RandomAI{generator: generator}
}

func (ai *RandomAI) ChooseMove(g *game.Game) int {
	moves := g.AvailableMoves()
	if len(moves) == 0 {
		return 0
	}
	return moves[ai.generator.NextValue(len(moves))]
}
