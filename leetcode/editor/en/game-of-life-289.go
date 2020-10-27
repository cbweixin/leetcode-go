package main

//According to the Wikipedia's article: "The Game of Life, also known simply as
//Life, is a cellular automaton devised by the British mathematician John Horton C
//onway in 1970."
//
// Given a board with m by n cells, each cell has an initial state live (1) or d
//ead (0). Each cell interacts with its eight neighbors (horizontal, vertical, dia
//gonal) using the following four rules (taken from the above Wikipedia article):
//
//
//
// Any live cell with fewer than two live neighbors dies, as if caused by under-
//population.
// Any live cell with two or three live neighbors lives on to the next generatio
//n.
// Any live cell with more than three live neighbors dies, as if by over-populat
//ion..
// Any dead cell with exactly three live neighbors becomes a live cell, as if by
// reproduction.
//
//
// Write a function to compute the next state (after one update) of the board gi
//ven its current state. The next state is created by applying the above rules sim
//ultaneously to every cell in the current state, where births and deaths occur si
//multaneously.
//
// Example:
//
//
//Input:
//[
//  [0,1,0],
//  [0,0,1],
//  [1,1,1],
//  [0,0,0]
//]
//Output:
//[
//  [0,0,0],
//  [1,0,1],
//  [0,1,1],
//  [0,1,0]
//]
//
//
// Follow up:
//
//
// Could you solve it in-place? Remember that the board needs to be updated at t
//he same time: You cannot update some cells first and then use their updated valu
//es to update other cells.
// In this question, we represent the board using a 2D array. In principle, the
//board is infinite, which would cause problems when the active area encroaches th
//e border of the array. How would you address these problems?
//
// Related Topics Array
// 👍 2090 👎 301

// 2020-10-26 13:24:53

//leetcode submit region begin(Prohibit modification and deletion)
func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	get_value := func(i, j int) int {
		value := 0
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := i+dx, j+dy
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					value += board[nx][ny] & 1
				}
			}
		}
		return value
	}

	// 00 -> 0 - 0 dead to dead
	// 01 -> 1 - 0, life to dead
	// 10 -> 0 -> 1, dead to life
	// 11 -> 1 -> 1 , life to life

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			value := get_value(i, j)
			if board[i][j] > 0 {
				if value < 2 || value > 3 {
					// 01, life to dead
					board[i][j] = 1
				} else {
					// 11 , life to life
					board[i][j] = 3
				}
			} else {
				if value == 3 {
					// 10 - dead to life
					board[i][j] = 2
				}
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
