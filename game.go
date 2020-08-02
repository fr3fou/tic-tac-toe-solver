package main

type Player int

const (
	// playerNil is the default value
	playerNil Player = iota
	PlayerX
	PlayerO
)

// Game represents a game of TicTacToe.
// `PlayerO` goes first.
type Game struct {
	State     [3][3]Player
	turnCount int
}

// NewGame is a game constructor.
func NewGame() *Game {
	return &Game{
		State: [3][3]Player{
			{playerNil, playerNil, playerNil},
			{playerNil, playerNil, playerNil},
			{playerNil, playerNil, playerNil},
		},
		turnCount: 0,
	}
}

// Put places the value on the coordinates.
func (g *Game) Put(i, j int) {
	switch g.turnCount % 2 {
	case 0:
		g.State[i][j] = PlayerO
	case 1:
		g.State[i][j] = PlayerX
	default:
		g.State[i][j] = playerNil
	}
	g.turnCount++
}
