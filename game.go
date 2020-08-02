package main

// Board is a game board.
// 0 means Blank.
// 1 means X.
// 2 means O.
type Board [3][3]int

// NewBoard is a constructor for a board>
func NewBoard() Board {
	return Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}
