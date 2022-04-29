package main

//给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
//
//
//
// 示例 1：
//
//
//输入：n = 13
//输出：6
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//
// Related Topics 递归 数学 动态规划 👍 416 👎 0

// 2022-04-29 16:39:55
//leetcode submit region begin(Prohibit modification and deletion)
func countDigitOne(n int) int {
	res := 0
	for k := 1; k <= n; k *= 10 {
		r, m := n%k, n/k
		res += (m + 8) / 10 * k
		if m%10 == 1 {
			res += (1 + r)
		}
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
