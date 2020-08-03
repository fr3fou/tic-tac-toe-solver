package main

func value(b Board, player Player) int {
	_, _ = b.Winner(player)
	return 0
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
