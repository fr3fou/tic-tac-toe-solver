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
	Board     Board
	turnCount int
	IsOver    bool
	Winner    Player
}

type Board [][]Player

// NewGame is a game constructor.
func NewGame() *Game {
	return &Game{
		Board: Board{
			{None, None, None},
			{None, None, None},
			{None, None, None},
		},
		IsOver:    false,
		Winner:    None,
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

			var color rl.Color
			switch {
			case g.Winner == None && g.IsOver:
				color = rl.Gray
			case (!g.IsOver && square == PlayerX) || g.Winner == PlayerX:
				color = rl.Blue
			case (!g.IsOver && square == PlayerO) || g.Winner == PlayerO:
				color = rl.Red
			}

			rl.DrawText(square.String(), x+CellSize/4, y-CellSize/20, CellSize, color)
		}
	}
}

// Update updates the state.
func (g *Game) Update() {
	g.IsOver, g.Winner = g.IsGameOver()
}

// IsGameOver checks if the game is over and returns the winning player if there is one.
// true and either `PlayerX` or `PlayerO` are returned if there is a winner.
// true and `None` are returned if there is a draw.
// false and `None` are returned if the game isn't over yet.
func (g *Game) IsGameOver() (bool, Player) {
	// Only the most recently played player can have made a winning placement
	recentPlayer := PlayerX
	if g.turnCount%2 != 0 {
		recentPlayer = PlayerO
	}

	// Horizontal check
	for _, row := range g.Board {
		if all(row[:], recentPlayer) {
			return true, recentPlayer
		}
	}

	// Vertical check
	for _, row := range transpose(g.Board) {
		if all(row[:], recentPlayer) {
			return true, recentPlayer
		}
	}

	// Primary diagonal
	wonDiagonally := true
	for i := 0; i < len(g.Board); i++ {
		if g.Board[i][i] != recentPlayer {
			wonDiagonally = false
			break
		}
	}
	if wonDiagonally {
		return true, recentPlayer
	}

	// Secondary diagonal
	wonDiagonally = true
	for i := len(g.Board) - 1; i >= 0; i-- {
		if g.Board[i][(Size-1)-i] != recentPlayer {
			wonDiagonally = false
			break
		}
	}
	if wonDiagonally {
		return true, recentPlayer
	}

	// Draw check
	if g.turnCount == Size*Size {
		return true, None
	}

	return false, None
}

// all checks if all the variables in the array are equal to the player passed.
func all(arr []Player, player Player) bool {
	for _, elem := range arr {
		if elem != player {
			return false
		}
	}

	return true
}

func transpose(arr Board) Board {
	transposed := Board{}

	for i, row := range arr {
		for j, square := range row {
			transposed[j][i] = square
		}
	}

	return transposed
}
