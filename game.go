package main

// Game represents a game of TicTacToe.
type Game struct {
	State     Board
	TurnCount int
}

// NewGame is a game constructor.
func NewGame() *Game {
	return &Game{
		State:     NewBoard(),
		TurnCount: 0,
	}
}

// Board is a game board.
// 0 means Blank.
// 1 means X.
// 2 means O.
type Board [3][3]int

// NewBoard is a constructor for a board.
func NewBoard() Board {
	return Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

// Put places the value on the coordinates
func (b Board) Put(i, j int, val int) {
	b[i][j] = val
}
