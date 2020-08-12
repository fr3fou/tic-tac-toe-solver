package main

import (
	"fmt"
	"strings"

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

type Board [][]Player

// IsWinner returns if the player passed is a winner or not
func (b Board) IsWinner(player Player) bool {
	// Only the most recently played player can have made a winning placement

	// Horizontal check
	for _, row := range b {
		if all(row[:], player) {
			return true
		}
	}

	// Vertical check
	for _, row := range transpose(b) {
		if all(row[:], player) {
			return true
		}
	}

	// Primary diagonal
	wonDiagonally := true
	for i := 0; i < len(b); i++ {
		if b[i][i] != player {
			wonDiagonally = false
			break
		}
	}
	if wonDiagonally {
		return true
	}

	// Secondary diagonal
	wonDiagonally = true
	for i := len(b) - 1; i >= 0; i-- {
		if b[i][(Size-1)-i] != player {
			wonDiagonally = false
			break
		}
	}
	if wonDiagonally {
		return true
	}

	return false
}

// Print prints the sudoku
func (board Board) String() string {
	b := &strings.Builder{}
	for _, line := range board {
		fmt.Fprintf(b, "+-----------+\n|")
		for _, num := range line {
			fmt.Fprintf(b, " %s |", num)
		}
		fmt.Fprintln(b)
	}
	fmt.Fprintln(b, "+-----------+")

	return b.String()
}

func (b Board) EmptySpots() int {
	n := 0
	for _, row := range b {
		for _, square := range row {
			if square == None {
				n++
			}
		}
	}
	return n
}

// Game represents a game of TicTacToe.
type Game struct {
	Board         Board
	IsOver        bool
	Winner        Player
	CurrentPlayer Player
}

// NewGame is a game constructor.
func NewGame(startingPlayer Player) *Game {
	return &Game{
		Board: Board{
			{None, None, None},
			{None, None, None},
			{None, None, None},
		},
		IsOver:        false,
		Winner:        None,
		CurrentPlayer: startingPlayer,
	}
}

func (g *Game) Place(i, j int) {
	// Prevent placing on occupied cells
	if g.Board[i][j] != None {
		return
	}

	g.Board[i][j] = g.CurrentPlayer
	g.CurrentPlayer = otherPlayer(g.CurrentPlayer)
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
	previousPlayer := otherPlayer(g.CurrentPlayer)

	didWin := g.Board.IsWinner(previousPlayer)
	isFull := g.Board.EmptySpots() == 0

	g.IsOver = didWin || isFull
	if didWin {
		g.Winner = previousPlayer
	}
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

func transpose(board Board) Board {
	transposed := Board{}

	// Init the board
	for i, row := range board {
		transposed = append(transposed, []Player{})
		for range row {
			transposed[i] = append(transposed[i], None)
		}
	}

	// Transpose
	for i, row := range board {
		for j, square := range row {
			transposed[j][i] = square
		}
	}

	return transposed
}

func otherPlayer(p Player) Player {
	switch p {
	case PlayerX:
		return PlayerO
	case PlayerO:
		return PlayerX
	default:
		return None
	}
}
