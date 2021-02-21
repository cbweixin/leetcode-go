package main

import (
	"fmt"
	"math"
	"sort"
)

//Given a non-empty 2D matrix matrix and an integer k, find the max sum of a rec
//tangle in the matrix such that its sum is no larger than k.
//
// Example:
//
//
//Input: matrix = [[1,0,1],[0,-2,3]], k = 2
//Output: 2
//Explanation: Because the sum of rectangle [[0, 1], [-2, 3]] is 2,
//             and 2 is the max number no larger than k (k = 2).
//
// Note:
//
//
// The rectangle inside the matrix must have an area > 0.
// What if the number of rows is much larger than the number of columns?
// Related Topics Binary Search Dynamic Programming Queue
// 👍 1033 👎 77

// 2021-02-21 10:12:06

//leetcode submit region begin(Prohibit modification and deletion)
func maxSumSubmatrix(matrix [][]int, k int) int {
	M, N := len(matrix), len(matrix[0])
	if M == 0 || N == 0 {
		return 0
	}

	fill := func(arr []int, val int) {
		for i := range arr {
			arr[i] = val
		}
	}

	max := func(x, y int) int {
		if x <= y {
			return y
		}

		return x
	}

	find_rect := func(arr []int) int {
		brr, ans := make([]int, 0, M), math.MinInt32
		brr = append(brr, 0)

		for i := 1; i < len(arr); i++ {
			arr[i] += arr[i-1]
		}

		for _, v := range arr {
			i := sort.Search(len(brr), func(idx int) bool { return v-brr[idx] <= k })
			if i < len(arr) && i >= 0 {
				ans = max(ans, v-brr[i])
			}

			i = sort.Search(len(brr), func(idx int) bool { return brr[idx] >= v })
			if i == len(brr) {
				brr = append(brr, v)
			} else {
				copy(brr[i+1:], brr[i:])
				brr[i] = v
			}

		}

		return ans

	}

	res := math.MinInt32
	for j := 0; j < N; j++ {
		row_sum := make([]int, M)
		fill(row_sum, 0)
		for i := j; i < N; i++ {
			for r := 0; r < M; r++ {
				row_sum[r] += matrix[r][i]

			}
			res = max(res, find_rect(row_sum))
		}

	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {

	matrix := [][]int{{2, 2, -1}}
	fmt.Println(maxSumSubmatrix(matrix, 0))
	//Input: matrix = [[1,0,1],[0,-2,3]], k = 2
	//Output: 2
	matrix = [][]int{{1, 0, 1}, {0, -2, 3}}
	fmt.Println(maxSumSubmatrix(matrix, 2))
	matrix = [][]int{{2, 2, -1}}
	fmt.Println(maxSumSubmatrix(matrix, 2))
}
