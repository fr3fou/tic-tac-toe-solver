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

// Game represents a game of TicTacToe.
// `PlayerO` goes first.
type Game struct {
	Board     [3][3]Player
	turnCount int
}

// NewGame is a game constructor.
func NewGame() *Game {
	return &Game{
		Board: [3][3]Player{
			{empty, empty, empty},
			{empty, empty, empty},
			{empty, empty, empty},
		},
		turnCount: 0,
	}
}

// Put places the value on the coordinates.
func (g *Game) Put(i, j int) {
	switch g.turnCount % 2 {
	case 0:
		g.Board[i][j] = PlayerO
	case 1:
		g.Board[i][j] = PlayerX
	default:
		g.Board[i][j] = empty
	}

	g.turnCount++
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.Gray)

	for i, row := range g.Board {
		for j := range row {
			rl.DrawRectangle(int32(i*CellSize+LineThickness*i), int32(j*CellSize+LineThickness*j), CellSize, CellSize, rl.Black)
		}
	}
}
