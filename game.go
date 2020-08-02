package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player int

const (
	// None is the default value
	None Player = iota
	PlayerX
	PlayerO
)

func (p Player) String() string {
	switch p {
	case PlayerO:
		return "o"
	case PlayerX:
		return "x"
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
			{None, None, None},
			{None, None, None},
			{None, None, None},
		},
		turnCount: 0,
	}
}

func (g *Game) Place(i, j int) {
	// Prevent placing on occupied cells
	if g.Board[i][j] != None {
		return
	}

	candidate := PlayerO
	if g.turnCount%2 != 0 {
		candidate = PlayerX
	}

	g.Board[i][j] = candidate
	g.turnCount++
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.Gray)

	for i, row := range transpose(g.Board) {
		for j, square := range row {
			x := int32(i*CellSize + LineThickness*i)
			y := int32(j*CellSize + LineThickness*j)
			rl.DrawRectangle(x, y, CellSize, CellSize, rl.Black)

			// No need to render empty text
			if square == None {
				continue
			}

			color := rl.Red
			if square == PlayerX {
				color = rl.Blue
			}

			rl.DrawText(square.String(), x+CellSize/4, y+CellSize/30, CellSize, color)
		}
	}
}

type State int

const (
	GameOver State = iota
	Draw
	XWins
	OWins
)

// func (g *Game) State() State {
// }

func transpose(arr [Size][Size]Player) [Size][Size]Player {
	t := [Size][Size]Player{}

	for i, row := range arr {
		for j, square := range row {
			t[j][i] = square
		}
	}

	return t
}
