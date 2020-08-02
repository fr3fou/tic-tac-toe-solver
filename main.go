package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Width         = 800
	Height        = 800
	LineThickness = 2
	CellSize      = (Width - LineThickness*2) / 3
)

func main() {
	rl.InitWindow(Width, Height, "Tic Tac Toe - AI")
	rl.SetTargetFPS(60)

	g := NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
