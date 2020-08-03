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

	g := NewGame(PlayerO)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		if !g.IsOver {
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) && g.CurrentPlayer == PlayerO {
				mousePos := rl.GetMousePosition()
				x := mousePos.X
				y := mousePos.Y
				i := int(x) / CellSize
				j := int(y) / CellSize
				// We swap i and j because the matrix is transposed
				g.Place(j, i)
			} else if g.CurrentPlayer == PlayerX {
				Minimax(PlayerX, g)
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
