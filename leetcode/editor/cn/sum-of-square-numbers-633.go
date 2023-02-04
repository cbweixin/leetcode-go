package main

import "math"

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a² + b² = c 。
//
//
//
// 示例 1：
//
//
//输入：c = 5
//输出：true
//解释：1 * 1 + 2 * 2 = 5
//
//
// 示例 2：
//
//
//输入：c = 3
//输出：false
//
//
//
//
// 提示：
//
//
// 0 <= c <= 2³¹ - 1
//
//
// Related Topics 数学 双指针 二分查找 👍 413 👎 0

//2023-02-04 10:05:49

// leetcode submit region begin(Prohibit modification and deletion)
func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))

	for a <= b {
		t := a*a + b*b
		if t == c {
			return true
		}
		if t < c {
			a++
		} else {
			b--
		}
	}
	return false

}

//leetcode submit region end(Prohibit modification and deletion)
