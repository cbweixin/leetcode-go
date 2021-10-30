package main

//给你一个整数数组 nums 和一个整数 k 。
//
// 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
//
// 返回你可以对数组执行的最大操作数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4], k = 5
//输出：2
//解释：开始时 nums = [1,2,3,4]：
//- 移出 1 和 4 ，之后 nums = [2,3]
//- 移出 2 和 3 ，之后 nums = []
//不再有和为 5 的数对，因此最多执行 2 次操作。
//
// 示例 2：
//
//
//输入：nums = [3,1,3,4,3], k = 6
//输出：1
//解释：开始时 nums = [3,1,3,4,3]：
//- 移出前两个 3 ，之后nums = [1,4,3]
//不再有和为 6 的数对，因此最多执行 1 次操作。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
// 1 <= k <= 10⁹
//
// Related Topics 数组 哈希表 双指针 排序 👍 19 👎 0

// 2021-10-30 09:18:00
//leetcode submit region begin(Prohibit modification and deletion)
func maxOperations(nums []int, k int) int {
	counter := make(map[int]int)
	for _, n := range nums {
		if _, ok := counter[n]; !ok {
			counter[n] = 1
		} else {
			counter[n]++
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	res := 0
	for key, v := range counter {
		if v2, ok := counter[k-key]; ok {
			if key<<1 == k {
				res += v / 2
				counter[k] = 0
			} else {
				n := min(v, v2)
				res += n
				counter[key] -= n
				counter[k-key] -= n
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
