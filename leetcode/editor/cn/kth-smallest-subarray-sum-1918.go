package main

import "math"

// 给你一个 长度为 n 的整型数组 nums 和一个数值 k ，返回 第 k 小的子数组和。
//
// 子数组 是指数组中一个 非空 且不间断的子序列。 子数组和 则指子数组中所有元素的和。
//
//
//
// 示例 1:
//
//
// 输入: nums = [2,1,3], k = 4
// 输出: 3
// 解释: [2,1,3] 的子数组为：
// - [2] 和为 2
// - [1] 和为 1
// - [3] 和为 3
// - [2,1] 和为 3
// - [1,3] 和为 4
// - [2,1,3] 和为 6
// 最小子数组和的升序排序为 1, 2, 3, 3, 4, 6。 第 4 小的子数组和为 3 。
//
//
// 示例 2：
//
//
// 输入：nums = [3,3,5,5], k = 7
// 输出：10
// 解释：[3,3,5,5] 的子数组为：
// - [3] 和为 3
// - [3] 和为 3
// - [5] 和为 5
// - [5] 和为 5
// - [3,3] 和为 6
// - [3,5] 和为 8
// - [5,5] 和为 10
// - [3,3,5], 和为 11
// - [3,5,5] 和为 13
// - [3,3,5,5] 和为 16
// 最小子数组和的升序排序为 3, 3, 5, 5, 6, 8, 10, 11, 13, 16。第 7 小的子数组和为 10 。
//
//
//
//
// 提示:
//
//
// n == nums.length
// 1 <= n <= 2 * 10⁴
// 1 <= nums[i] <= 5 * 10⁴
// 1 <= k <= n * (n + 1) / 2
//
//
// Related Topics 数组 二分查找 滑动窗口 👍 16 👎 0

// 2022-12-31 16:25:12
// leetcode submit region begin(Prohibit modification and deletion)
func kthSmallestSubarraySum(nums []int, k int) int {

	countSmall := func(n int) int {
		l, r, s, cnt := 0, 0, 0, 0
		for r < len(nums) {
			s += nums[r]
			for s > n {
				s -= nums[l]
				l++
			}
			cnt += r - l + 1
			r++
		}
		return cnt
	}

	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	minN, s := math.MaxInt, 0
	for _, n := range nums {
		minN = min(minN, n)
		s += n
	}

	l, h := 0, s

	for l < h {
		mid := (l + h) >> 1
		if countSmall(mid) < k {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}

// leetcode submit region end(Prohibit modification and deletion)
