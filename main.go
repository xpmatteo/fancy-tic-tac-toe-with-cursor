//go:build js && wasm

package main

import (
	"syscall/js"
	"tictactoe/game"
	"tictactoe/perfectai"
)

func main() {
	c := make(chan struct{}, 0)

	g := game.NewGame()
	ai := perfectai.NewPerfectAI()

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
			if err := g.MakeMove(position); err != nil {
				return err.Error()
			}
			if g.Winner() == "" {
				// AI's turn
				aiMove := ai.ChooseMove(g)
				if err := g.MakeMove(aiMove); err != nil {
					return err.Error()
				}
			}
			return ""
		}),
		"winner": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return g.Winner()
		}),
	}))

	<-c // Keep running
}
