package main

//Alex and Lee play a game with piles of stones. There are an even number of pil
//es arranged in a row, and each pile has a positive integer number of stones pile
//s[i].
//
// The objective of the game is to end with the most stones. The total number of
// stones is odd, so there are no ties.
//
// Alex and Lee take turns, with Alex starting first. Each turn, a player takes
//the entire pile of stones from either the beginning or the end of the row. This
//continues until there are no more piles left, at which point the person with the
// most stones wins.
//
// Assuming Alex and Lee play optimally, return True if and only if Alex wins th
//e game.
//
//
// Example 1:
//
//
//Input: piles = [5,3,4,5]
//Output: true
//Explanation:
//Alex starts first, and can only take the first 5 or the last 5.
//Say he takes the first 5, so that the row becomes [3, 4, 5].
//If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 poin
//ts.
//If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win wit
//h 9 points.
//This demonstrated that taking the first 5 was a winning move for Alex, so we r
//eturn true.
//
//
//
// Constraints:
//
//
// 2 <= piles.length <= 500
// piles.length is even.
// 1 <= piles[i] <= 500
// sum(piles) is odd.
//
// Related Topics Math Dynamic Programming Minimax
// 👍 984 👎 1239

// 2021-02-05 09:53:07

//leetcode submit region begin(Prohibit modification and deletion)
func stoneGame(piles []int) bool {
	l := len(piles)

	// dp[i][j], max diff of stones when pick from i -> j
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	for i := 0; i < l; i++ {
		dp[i][i] = piles[i]
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	// d , distance between i, j
	for d := 1; d < l; d++ {
		for i := 0; i < l-d; i++ {
			dp[i][i+d] = max(piles[i]-dp[i+1][i+d], piles[i+d]-dp[i][i+d-1])
		}
	}

	return dp[0][l-1] > 0

}

//leetcode submit region end(Prohibit modification and deletion)
