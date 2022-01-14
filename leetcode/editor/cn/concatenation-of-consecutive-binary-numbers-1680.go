package main

import "math/bits"

//给你一个整数 n ，请你将 1 到 n 的二进制表示连接起来，并返回连接结果对应的 十进制 数字对 10⁹ + 7 取余的结果。
//
//
//
// 示例 1：
//
// 输入：n = 1
//输出：1
//解释：二进制的 "1" 对应着十进制的 1 。
//
//
// 示例 2：
//
// 输入：n = 3
//输出：27
//解释：二进制下，1，2 和 3 分别对应 "1" ，"10" 和 "11" 。
//将它们依次连接，我们得到 "11011" ，对应着十进制的 27 。
//
//
// 示例 3：
//
// 输入：n = 12
//输出：505379714
//解释：连接结果为 "1101110010111011110001001101010111100" 。
//对应的十进制数字为 118505380540 。
//对 10⁹ + 7 取余后，结果为 505379714 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
//
// Related Topics 位运算 数学 模拟 👍 26 👎 0

// 2021-10-30 11:06:36
//leetcode submit region begin(Prohibit modification and deletion)
// bit_manipulation
func concatenatedBinary_2(n int) (ans int) {
	shift, ans := 0, 0
	mod := int(1e9 + 7)

	for i := 1; i < n+1; i++ {
		if i&(i-1) == 0 {
			shift++
		}
		ans = ((ans << shift) + i) % mod
	}
	return
}

func concatenatedBinary(n int) (ans int) {
	for i := 1; i < n+1; i++ {
		ans = (ans<<bits.Len(uint(i)) | i) % int(1e9+7)
	}
	return

}

//leetcode submit region end(Prohibit modification and deletion)
