package main

import (
	"math/rand"
	"time"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2164 👎 0

//2023-05-08 19:18:41

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {

	partiton := func(l, r, pivot int) int {
		nums[r], nums[pivot] = nums[pivot], nums[r]
		k := l

		for i := l; i < r; i++ {
			if nums[i] > nums[r] {
				nums[i], nums[k] = nums[k], nums[i]
				k += 1
			}
		}
		nums[k], nums[r] = nums[r], nums[k]

		return k
	}

	lo, hi := 0, len(nums)-1

	for lo <= hi {
		rand.Seed(time.Now().UnixNano())
		mid := rand.Intn(hi-lo+1) + lo
		r := partiton(lo, hi, mid)
		if r == k-1 {
			return nums[r]
		}
		if r > k-1 {
			hi = r - 1
		} else {
			lo = r + 1
		}
	}

	return nums[lo]
}

//leetcode submit region end(Prohibit modification and deletion)
