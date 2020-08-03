package main

func nextBoards(b Board, player Player) []Board {
	boards := []Board{}

	for i, row := range b {
		for j, square := range row {
			if square != None {
				continue
			}

			var board Board
			copy(board, b)

			board[i][j] = player
			boards = append(boards, board)
		}
	}

	return boards
}
