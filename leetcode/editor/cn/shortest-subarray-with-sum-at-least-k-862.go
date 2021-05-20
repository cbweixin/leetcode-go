package main

//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
//
// 如果没有和至少为 K 的非空子数组，返回 -1 。
//
//
//
//
//
//
// 示例 1：
//
// 输入：A = [1], K = 1
//输出：1
//
//
// 示例 2：
//
// 输入：A = [1,2], K = 4
//输出：-1
//
//
// 示例 3：
//
// 输入：A = [2,-1,2], K = 3
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 50000
// -10 ^ 5 <= A[i] <= 10 ^ 5
// 1 <= K <= 10 ^ 9
//
// Related Topics 队列 二分查找
// 👍 278 👎 0

// 2021-05-18 21:29:26
//leetcode submit region begin(Prohibit modification and deletion)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func shortestSubarray(nums []int, k int) int {
	l := len(nums)
	preSum := make([]int, l+1)
	for i := 0; i < l; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	// review how to simulate a stack
	queue := make([]int, 0)
	res := l + 1
	for i, v := range preSum {
		for len(queue) > 0 && v <= preSum[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		for len(queue) > 0 && v-preSum[queue[0]] >= k {
			res = min(res, i-queue[0])
			queue = queue[1:]
		}
		queue = append(queue, i)
	}

	if res == l+1 {
		return -1
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
