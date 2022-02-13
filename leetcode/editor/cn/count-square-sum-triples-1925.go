package main

import "math"

//一个 平方和三元组 (a,b,c) 指的是满足 a² + b² = c² 的 整数 三元组 a，b 和 c 。
//
// 给你一个整数 n ，请你返回满足 1 <= a, b, c <= n 的 平方和三元组 的数目。
//
//
//
// 示例 1：
//
// 输入：n = 5
//输出：2
//解释：平方和三元组为 (3,4,5) 和 (4,3,5) 。
//
//
// 示例 2：
//
// 输入：n = 10
//输出：4
//解释：平方和三元组为 (3,4,5)，(4,3,5)，(6,8,10) 和 (8,6,10) 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 250
//
// Related Topics 数学 枚举 👍 9 👎 0

// 2022-02-13 12:39:46
//leetcode submit region begin(Prohibit modification and deletion)
func countTriples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			sq := i*i + j*j
			r := int(math.Sqrt(float64(i*i + j*j)))
			if r*r == sq && r <= n {
				res += 2
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
