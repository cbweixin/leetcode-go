package main

import (
	"sort"
)

//Given an array of integers A with even length, return true if and only if it i
//s possible to reorder it such that A[2 * i + 1] = 2 * A[2 * i] for every 0 <= i
//< len(A) / 2.
//
//
//
//
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
//Input: A = [3,1,3,6]
//Output: false
//
//
// Example 2:
//
//
//Input: A = [2,1,2,6]
//Output: false
//
//
// Example 3:
//
//
//Input: A = [4,-2,2,-4]
//Output: true
//Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or
//[2,4,-2,-4].
//
//
// Example 4:
//
//
//Input: A = [1,2,4,16,8,4]
//Output: false
//
//
//
// Constraints:
//
//
// 0 <= A.length <= 3 * 104
// A.length is even.
// -105 <= A[i] <= 105
//
// Related Topics Array Hash Table
// 👍 323 👎 55

// 2020-10-25 17:39:37

//leetcode submit region begin(Prohibit modification and deletion)
func canReorderDoubled(A []int) bool {
	ans := []int{}
	sort.Ints(A)
	dict := map[int]int{}
	for _, n := range A {
		if _, ok := dict[n]; ok {
			dict[n]++
		} else {
			dict[n] = 1
		}
	}

	for _, n := range A {
		if count, ok := dict[n]; ok && count > 0 {
			k := n << 1
			if count, ok := dict[k]; ok && count > 0 {
				dict[n] -= 1
				dict[k] -= 1
				ans = append(ans, n)
				ans = append(ans, k)
			}
		}
	}

	return len(ans) == len(A)
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	println(canReorderDoubled([]int{4, -2, 2, -4}))

}
