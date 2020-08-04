package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	LineThickness = 2
	Width         = 800 + LineThickness
	Height        = 800 + LineThickness
	CellSize      = (Width - LineThickness*2) / 3
)

func main() {
	rl.InitWindow(Width, Height, "Tic Tac Toe - AI")
	rl.SetTargetFPS(60)

	user := PlayerO
	ai := PlayerX
	g := NewGame(PlayerO)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		if !g.IsOver {
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) && g.CurrentPlayer == user {
				mousePos := rl.GetMousePosition()
				x := mousePos.X
				y := mousePos.Y
				i := int(x) / CellSize
				j := int(y) / CellSize
				// We swap i and j because the matrix is transposed
				g.Place(j, i)
			} else if g.CurrentPlayer == ai {
				Minimax(ai, g)
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
