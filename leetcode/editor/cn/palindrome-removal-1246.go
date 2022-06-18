package main

import "math"

//给你一个整数数组 arr，每一次操作你都可以选择并删除它的一个 回文 子数组 arr[i], arr[i+1], ..., arr[j]（ i <= j）。
//
//
// 注意，每当你删除掉一个子数组，右侧元素都会自行向前移动填补空位。
//
// 请你计算并返回从数组中删除所有数字所需的最少操作次数。
//
//
//
// 示例 1：
//
// 输入：arr = [1,2]
//输出：2
//
//
// 示例 2：
//
// 输入：arr = [1,3,4,1,5]
//输出：3
//解释：先删除 [4]，然后删除 [1,3,1]，最后再删除 [5]。
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 100
// 1 <= arr[i] <= 20
//
// Related Topics 数组 动态规划 👍 65 👎 0

// 2022-06-18 11:26:00
//leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(arr []int) int {
	l := len(arr)
	dp := make([][]int, l)
	for i, _ := range dp {
		dp[i] = make([]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 0; i < l; i++ {
		dp[i][i] = 1
	}
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	for gap := 1; gap < l; gap++ {
		for i := l - 1 - gap; i >= 0; i-- {
			j := i + gap
			if arr[i] == arr[j] {
				if gap == 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
		}
	}

	return dp[0][l-1]

}

//leetcode submit region end(Prohibit modification and deletion)
