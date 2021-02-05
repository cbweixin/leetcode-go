package main

import "math"

//Given N, consider a convex N-sided polygon with vertices labelled A[0], A[i],
//..., A[N-1] in clockwise order.
//
// Suppose you triangulate the polygon into N-2 triangles. For each triangle, th
//e value of that triangle is the product of the labels of the vertices, and the t
//otal score of the triangulation is the sum of these values over all N-2 triangle
//s in the triangulation.
//
// Return the smallest possible total score that you can achieve with some trian
//gulation of the polygon.
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
//Input: [1,2,3]
//Output: 6
//Explanation: The polygon is already triangulated, and the score of the only tr
//iangle is 6.
//
//
//
// Example 2:
//
//
//
//
//Input: [3,7,4,5]
//Output: 144
//Explanation: There are two triangulations, with possible scores: 3*7*5 + 4*5*7
// = 245, or 3*4*5 + 3*4*7 = 144.  The minimum score is 144.
//
//
//
// Example 3:
//
//
//Input: [1,3,1,4,1,5]
//Output: 13
//Explanation: The minimum score triangulation has score 1*1*3 + 1*1*4 + 1*1*5 +
// 1*1*1 = 13.
//
//
//
//
// Note:
//
//
// 3 <= A.length <= 50
// 1 <= A[i] <= 100
//
//
//
//
// Related Topics Dynamic Programming
// 👍 585 👎 77

// 2021-02-05 12:53:27

//leetcode submit region begin(Prohibit modification and deletion)
func minScoreTriangulation(A []int) int {
	l := len(A)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	for d := 2; d < l; d++ {
		for i := 0; i < l-d; i++ {
			j := i + d
			t := math.MaxInt32
			for k := i + 1; k < j; k++ {
				t = min(t, dp[i][k]+A[i]*A[k]*A[j]+dp[k][j])
			}
			dp[i][j] = t
		}
	}

	return dp[0][l-1]

}

//leetcode submit region end(Prohibit modification and deletion)
