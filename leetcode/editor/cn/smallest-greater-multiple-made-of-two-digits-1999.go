package main

import (
	"fmt"
	"math"
)

//给你三个整数, k, digit1和 digit2, 你想要找到满足以下条件的 最小 整数：
//
//
// 大于k 且是 k 的倍数
// 仅由digit1 和 digit2 组成，即 每一位数 均是 digit1 或 digit2
//
//
// 请你返回 最小的满足这两个条件的整数，如果不存在这样的整数，或者最小的满足这两个条件的整数不在32位整数范围（0~2³¹-1），就返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：k = 2, digit1 = 0, digit2 = 2
//输出：20
//解释：
//20 是第一个仅有数字0和2组成的，比2大且是2的倍数的整数。
//
//
// 示例 2：
//
//
//输入：k = 3, digit1 = 4, digit2 = 2
//输出：24
//解释：
//24 是第一个仅有数字 2 和 4 组成的，比 3 大且是 3 的倍数的整数。
//
// 示例 3：
//
//
//输入：k = 2, digit1 = 0, digit2 = 0
//输出：-1
//解释：
//不存在仅由 0 组成的比 2 大且是 2 的倍数的整数，因此返回 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= k <= 1000
// 0 <= digit1 <= 9
// 0 <= digit2 <= 9
//
//
// Related Topics 数学 枚举 👍 6 👎 0

//2023-01-26 13:51:34

// leetcode submit region begin(Prohibit modification and deletion)
func findInteger2(k int, digit1 int, digit2 int) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	var search func(int) int
	search = func(i int) int {
		if i > math.MaxInt32 {
			return -1
		}

		if i > k && i%k == 0 {
			return i
		}

		a := -1
		if i+digit1 > 0 {
			a = search(i*10 + digit1)
		}

		b := -1
		if i+digit2 > 0 {
			b = search(i*10 + digit2)
		}

		if a == -1 || b == -1 {
			return max(a, b)
		}

		return min(a, b)
	}

	return search(0)

}

func findInteger(k int, digit1 int, digit2 int) int {
	if digit1+digit2 == 0 {
		return -1
	}
	if digit1 > digit2 {
		digit1, digit2 = digit2, digit1
	}
	var deque = []int{digit1, digit2}
	for len(deque) > 0 {
		a := deque[0]
		if a > k && a%k == 0 {
			return a
		}
		deque = append(deque[1:], []int{a*10 + digit1, a*10 + digit2}...)
	}

	return -1

}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findInteger(3, 3, 6))
}
