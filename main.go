//go:build js && wasm

package main

import (
	"syscall/js"
	"tictactoe/game"
	"tictactoe/randomai"
)

func main() {
	c := make(chan struct{}, 0)

	g := game.NewGame()
	ai := randomai.NewRandomAI(randomai.NewRealRNG())

	// Register our game in JavaScript
	js.Global().Set("game", js.ValueOf(map[string]interface{}{
		"text": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return g.Text()
		}),
		"board": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			board := g.Board()
			// Convert Go array to JavaScript array
			jsBoard := make([]interface{}, len(board))
			for i, cell := range board {
				jsBoard[i] = cell
			}
			return jsBoard
		}),
		"makeMove": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			position := args[0].Int()
			err := g.MakeMove(position)
			if err == nil && g.Winner() == "" {
				// AI's turn
				aiMove := ai.ChooseMove(g)
				g.MakeMove(aiMove)
			}
			return nil
		}),
		"winner": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return g.Winner()
		}),
	}))

	<-c // Keep running
}
