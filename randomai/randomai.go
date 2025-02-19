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
	return 0 // This will make the test fail as it always returns 0
}
