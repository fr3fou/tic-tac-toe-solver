package main

import "math"

// value should only be called at leaf / terminal nodes (game MUST be over).
func value(b Board, aiWon bool) int {
	spots := b.EmptySpots()

	if aiWon {
		return max(1, spots)
	}

	return min(-1, -spots)
}

func minimax(b Board, ai, current Player, alpha, beta float64) int {
	other := otherPlayer(current)

	if b.IsWinner(other) {
		return value(b, ai == other)
	}

	// Draw is 0
	if b.EmptySpots() == 0 {
		return 0
	}

	if current == ai {
		// Maximizing
		max := math.Inf(-1)
		for _, state := range nextBoards(b, current) {
			value := float64(minimax(state.Board, ai, other, alpha, beta))
			max = math.Max(max, value)
			alpha = math.Max(max, alpha)
			if alpha >= beta {
				break
			}
		}
		return int(max)
	} else {
		// Minimizing
		min := math.Inf(+1)
		for _, state := range nextBoards(b, current) {
			value := float64(minimax(state.Board, ai, other, alpha, beta))
			min = math.Min(min, value)
			beta = math.Min(min, beta)
			if alpha >= beta {
				break
			}
		}
		return int(min)
	}
}

func Minimax(ai Player, g *Game) {
	max := math.Inf(-1)
	bestState := State{}
	other := otherPlayer(ai)
	for _, state := range nextBoards(g.Board, ai) {
		value := float64(minimax(state.Board, ai, other, math.Inf(-1), math.Inf(1)))
		if value > max {
			max = value
			bestState = state
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
