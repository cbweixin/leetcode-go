package main

import "fmt"

//There is an m by n grid with a ball. Given the start coordinate (i,j) of the b
//all, you can move the ball to adjacent cell or cross the grid boundary in four d
//irections (up, down, left, right). However, you can at most move N times. Find o
//ut the number of paths to move the ball out of grid boundary. The answer may be
//very large, return it after mod 109 + 7.
//
//
//
// Example 1:
//
//
//Input: m = 2, n = 2, N = 2, i = 0, j = 0
//Output: 6
//Explanation:
//
//
//
// Example 2:
//
//
//Input: m = 1, n = 3, N = 3, i = 0, j = 1
//Output: 12
//Explanation:
//
//
//
//
//
// Note:
//
//
// Once you move the ball out of boundary, you cannot move it back.
// The length and height of the grid is in range [1,50].
// N is in range [0,50].
//
// Related Topics Dynamic Programming Depth-first Search
// 👍 731 👎 148

// 2021-03-13 09:45:54

//leetcode submit region begin(Prohibit modification and deletion)
func findPaths(m int, n int, N int, i int, j int) int {
	res := 0
	dp := make([][]int, m)
	for x := range dp {
		dp[x] = make([]int, n)
		for y := 0; y < n; y++ {
			dp[x][y] = 0
		}
	}

	dp[i][j] = 1
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	MOD := 1000000007

	for k := 0; k < N; k++ {
		temp := make([][]int, m)
		for x := range temp {
			temp[x] = make([]int, n)
		}
		for a := 0; a < m; a++ {
			for b := 0; b < n; b++ {
				for _, dir := range dirs {
					n_a, n_b := a+dir[0], b+dir[1]
					if n_a < 0 || n_a >= m || n_b < 0 || n_b >= n {
						res = (res + dp[a][b]) % MOD
					} else {
						temp[n_a][n_b] = (temp[n_a][n_b] + dp[a][b]) % MOD
					}
				}
			}
			dp = temp
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
}
