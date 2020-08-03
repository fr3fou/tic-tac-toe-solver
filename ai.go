package main

import (
	"math"
)

// value should only be called at leaf / terminal nodes (game MUST be over).
func value(b Board, player, winner Player) int {
	// Draws are 0
	if winner == None {
		return 0
	}

	empty := b.EmptySpots()
	if winner != player {
		// Prevent 0 value
		return min(-1, -empty)
	}

	// Prevent 0 value
	return max(1, empty)
}

func minimax(b Board, ai, current Player) int {
	// Terminal node
	if isOver, winner := b.Winner(current); isOver {
		return value(b, current, winner)
	}

	var other Player
	switch current {
	case PlayerX:
		other = PlayerO
	case PlayerO:
		other = PlayerX
	}

	if current == ai {
		// Maximizing
		max := math.Inf(-1)
		for _, state := range nextBoards(b, current) {
			value := minimax(state.Board, ai, other)
			max = math.Max(max, float64(value))
		}
		return int(max)
	} else {
		// Minimizing
		min := math.Inf(1)
		for _, state := range nextBoards(b, current) {
			value := minimax(state.Board, ai, other)
			min = math.Min(min, float64(value))
		}
		return int(min)
	}
}

func Minimax(ai Player, g *Game) {
	max := math.Inf(-1)
	bestState := State{}
	for _, state := range nextBoards(g.Board, ai) {
		value := float64(minimax(state.Board, ai, ai))
		if value > max {
			bestState = state
			max = value
		}
	}

	g.Place(bestState.X, bestState.Y)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type State struct {
	X     int
	Y     int
	Board Board
}

func nextBoards(b Board, player Player) []State {
	boards := []State{}

	for i, row := range b {
		for j, square := range row {
			if square != None {
				continue
			}

			board := copyBoard(b)

			board[i][j] = player
			boards = append(boards, State{
				X:     i,
				Y:     j,
				Board: board,
			})
		}
	}

	return boards
}

func copyBoard(src Board) Board {
	dst := Board{}

	for i, row := range src {
		dst = append(dst, []Player{})
		for _, square := range row {
			dst[i] = append(dst[i], square)
		}
	}

	return dst
}
