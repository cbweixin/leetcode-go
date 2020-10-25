package main

//We have a two dimensional matrix A where each value is 0 or 1.
//
// A move consists of choosing any row or column, and toggling each value in tha
//t row or column: changing all 0s to 1s, and all 1s to 0s.
//
// After making any number of moves, every row of this matrix is interpreted as
//a binary number, and the score of the matrix is the sum of these numbers.
//
// Return the highest possible score.
//
//
//
//
//
//
//
// Example 1:
//
//
//Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
//Output: 39
//Explanation:
//Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
//0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
//
//
//
// Note:
//
//
// 1 <= A.length <= 20
// 1 <= A[0].length <= 20
// A[i][j] is 0 or 1.
//
//
// Related Topics Greedy
// 👍 549 👎 122

// 2020-10-25 10:50:35

//leetcode submit region begin(Prohibit modification and deletion)
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func matrixScore(A [][]int) int {
	M, N := len(A), len(A[0])
	// review, need to convert int to unsigned int
	ans := (1 << uint(N-1)) * M

	for j := 1; j < N; j++ {
		cnt := 0
		for i := 0; i < M; i++ {
			if A[i][j] == A[i][0] {
				cnt++
			}
		}
		ans += max(cnt, M-cnt) * (1 << uint(N-j-1))
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
