package main

import (
	"fmt"
	"sort"
)

//You are given an m * n matrix, mat, and an integer k, which has its rows sorte
//d in non-decreasing order.
//
// You are allowed to choose exactly 1 element from each row to form an array. R
//eturn the Kth smallest array sum among all possible arrays.
//
//
// Example 1:
//
//
//Input: mat = [[1,3,11],[2,4,6]], k = 5
//Output: 7
//Explanation: Choosing one element from each row, the first k smallest sum are:
//
//[1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.
//
// Example 2:
//
//
//Input: mat = [[1,3,11],[2,4,6]], k = 9
//Output: 17
//
//
// Example 3:
//
//
//Input: mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
//Output: 9
//Explanation: Choosing one element from each row, the first k smallest sum are:
//
//[1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]. Where the 7th s
//um is 9.
//
//
// Example 4:
//
//
//Input: mat = [[1,1,10],[2,2,9]], k = 7
//Output: 12
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat.length[i]
// 1 <= m, n <= 40
// 1 <= k <= min(200, n ^ m)
// 1 <= mat[i][j] <= 5000
// mat[i] is a non decreasing array.
//
// Related Topics Heap
// 👍 446 👎 7

// 2021-02-15 10:59:12

//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallest(mat [][]int, k int) int {
	m := len(mat)
	// knowledge, how to copy an array
	temp := mat[0]
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < m; i++ {
		t := []int{}
		for _, v := range mat[i] {
			for _, v2 := range temp {
				t = append(t, v+v2)
			}
		}
		sort.Ints(t)
		temp = t[:min(k, len(t))]
	}

	return temp[k-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	// knowledge, how to construct and intialize an multi-dimensional array
	var a = [][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}
	fmt.Println(kthSmallest(a, 14))
}
