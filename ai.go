package main

// value should only be called at leaf / terminal nodes (game MUST be over).
func value(b Board, player Player) int {
	isOver, winner := b.Winner(player)

	// Draws are 0
	if winner == None && isOver {
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
