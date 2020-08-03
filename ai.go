package main

import "math"

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

func Minimax(ai Player) func(b Board, player Player) int {
	var minimax func(b Board, player Player) int
	minimax = func(b Board, player Player) int {
		// Terminal node
		if isOver, winner := b.Winner(player); isOver {
			return value(b, player, winner)
		}
		var other Player
		switch player {
		case PlayerX:
			other = PlayerO
		case PlayerO:
			other = PlayerX
		}

		// Maximizing
		if player == ai {
			max := math.Inf(-1)
			for _, state := range nextBoards(b, player) {
				value := minimax(state, other)
				max = math.Max(max, float64(value))
			}
			return int(max)
		} else {
			min := math.Inf(1)
			for _, state := range nextBoards(b, player) {
				value := minimax(state, other)
				min = math.Min(min, float64(value))
			}
			return int(min)
		}
	}

	return minimax
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

func nextBoards(b Board, player Player) []Board {
	boards := []Board{}

	for i, row := range b {
		for j, square := range row {
			if square != None {
				continue
			}

			board := copyBoard(b)

			board[i][j] = player
			boards = append(boards, board)
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
