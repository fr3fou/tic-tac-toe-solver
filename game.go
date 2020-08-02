package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player int

const (
	// empty is the default value
	empty Player = iota
	PlayerX
	PlayerO
)

func (p Player) String() string {
	switch p {
	case PlayerO:
		return "O"
	case PlayerX:
		return "X"
	default:
		return " "
	}
}

const Size = 3

// Game represents a game of TicTacToe.
// `PlayerO` goes first.
type Game struct {
	Board     [Size][Size]Player
	turnCount int
}

// NewGame is a game constructor.
func NewGame() *Game {
	return &Game{
		Board: [Size][Size]Player{
			{empty, PlayerX, PlayerO},
			{PlayerX, PlayerO, PlayerX},
			{empty, PlayerX, PlayerO},
		},
		turnCount: 0,
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.Gray)

	for i, row := range g.Board {
		for j, square := range row {
			x := int32(i*CellSize + LineThickness*i)
			y := int32(j*CellSize + LineThickness*j)
			rl.DrawRectangle(x, y, CellSize, CellSize, rl.Black)

			// No need to render empty text
			if square == empty {
				continue
			}

			color := rl.Red
			if square == PlayerX {
				color = rl.Blue
			}

			rl.DrawText(square.String(), x+CellSize/5, y+CellSize/20, CellSize, color)
		}
	}
}
