package main

type Key struct {
	X, Y int
}

func candyCrush(board [][]int) [][]int {
	m, n := len(board), len(board[0])
	for {
		// knowledge, how to create a composite key
		crush := map[Key]bool{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == 0 {
					continue
				}
				if i > 1 && board[i][j] == board[i-1][j] && board[i][j] == board[i-2][j] {
					crush[Key{i, j}] = true
					crush[Key{i - 1, j}] = true
					crush[Key{i - 2, j}] = true

				}
				if j > 1 && board[i][j] == board[i][j-1] && board[i][j] == board[i][j-2] {
					crush[Key{i, j}] = true
					crush[Key{i, j - 1}] = true
					crush[Key{i, j - 2}] = true
				}
			}
		}

		if len(crush) == 0 {
			break
		}

		for K := range crush {
			board[K.X][K.Y] = 0
		}

		for j := 0; j < n; j++ {
			bottom := m - 1
			for i := m - 1; i >= 0; i-- {
				if board[i][j] > 0 {
					board[bottom][j] = board[i][j]
					bottom--
				}

			}

			for i := 0; i <= bottom; i++ {
				board[i][j] = 0
			}
		}

	}

	return board

}
