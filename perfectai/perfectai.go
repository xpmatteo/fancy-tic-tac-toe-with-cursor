package perfectai

import "tictactoe/game"

type PerfectAI struct{}

func NewPerfectAI() *PerfectAI {
	return &PerfectAI{}
}

func (ai *PerfectAI) ChooseMove(g *game.Game) int {
	return 0 // This will make the test fail
}
