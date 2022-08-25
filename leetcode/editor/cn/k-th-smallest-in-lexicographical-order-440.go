package main

// 给定整数 n 和 k，返回 [1, n] 中字典序第 k 小的数字。
//
//
//
// 示例 1:
//
//
// 输入: n = 13, k = 2
// 输出: 10
// 解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//
//
// 示例 2:
//
//
// 输入: n = 1, k = 1
// 输出: 1
//
//
//
//
// 提示:
//
//
// 1 <= k <= n <= 10⁹
//
//
// Related Topics 字典树 👍 527 👎 0

// 2022-08-24 22:42:11
// leetcode submit region begin(Prohibit modification and deletion)
func findKthNumber(n int, k int) int {
	cur, k := 1, k-1
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	findSteps := func(a, cur, next int) int {
		steps := 0
		for cur <= a {
			steps += min(a+1, next) - cur
			cur, next = cur*10, next*10
		}
		return steps
	}

	for k > 0 {
		steps := findSteps(n, cur, cur+1)
		if steps <= k {
			k -= steps
			cur = cur + 1
		} else {
			k -= 1
			cur *= 10
		}
	}

	return cur

}

// leetcode submit region end(Prohibit modification and deletion)
