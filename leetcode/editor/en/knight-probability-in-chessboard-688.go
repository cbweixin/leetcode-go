package main

import "fmt"

//On an NxN chessboard, a knight starts at the r-th row and c-th column and atte
//mpts to make exactly K moves. The rows and columns are 0 indexed, so the top-lef
//t square is (0, 0), and the bottom-right square is (N-1, N-1).
//
// A chess knight has 8 possible moves it can make, as illustrated below. Each m
//ove is two squares in a cardinal direction, then one square in an orthogonal dir
//ection.
//
//
//
//
//
//
//
// Each time the knight is to move, it chooses one of eight possible moves unifo
//rmly at random (even if the piece would go off the chessboard) and moves there.
//
//
// The knight continues moving until it has made exactly K moves or has moved of
//f the chessboard. Return the probability that the knight remains on the board af
//ter it has stopped moving.
//
//
//
// Example:
//
//
//Input: 3, 2, 0, 0
//Output: 0.0625
//Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight o
//n the board.
//From each of those positions, there are also two moves that will keep the knig
//ht on the board.
//The total probability the knight stays on the board is 0.0625.
//
//
//
//
// Note:
//
//
// N will be between 1 and 25.
// K will be between 0 and 100.
// The knight always initially starts on the board.
//
// Related Topics Dynamic Programming
// 👍 1198 👎 225

// 2021-03-10 19:31:30

//leetcode submit region begin(Prohibit modification and deletion)
type Key1 struct {
	K, R, C int
}

func knightProbability(N int, K int, r int, c int) float64 {
	dirs := [][]int{{-2, -1}, {-1, -2}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {1, -2}, {2, -1}}
	probs := make(map[Key1]float64)

	isValid := func(row, col int) bool {
		return row < N && row >= 0 && col < N && col >= 0
	}

	var dfs func(step, row, col int, p float64) float64
	dfs = func(step, row, col int, p float64) float64 {
		if isValid(row, col) && step < K {
			ps := float64(0)
			for _, dir := range dirs {
				nR, nC := row+dir[0], col+dir[1]
				key := Key1{step, nR, nC}
				if isValid(nR, nC) {
					if _, ok := probs[key]; !ok {
						probs[key] = dfs(step+1, nR, nC, float64(p/8))
					}
				}
				ps += probs[key]
			}
			probs[Key1{step, row, col}] = ps
			return ps
		} else {
			if isValid(row, col) {
				return p
			}
		}
		return 0
	}

	return dfs(0, r, c, 1)

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))

}
