package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 4072 👎 0

//2023-02-04 15:14:23

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	n := len(height)
	l, r := 0, n-1
	res := 0
	for l < r {
		if height[l] < height[r] {
			k := l + 1
			for k < r && height[k] < height[l] {
				res += height[l] - height[k]
				k++
			}
			l = k
		} else {
			k := r - 1
			for k > l && height[k] < height[r] {
				res += height[r] - height[k]
				k--
			}
			r = k
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
