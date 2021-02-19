package main

//Given a sorted positive integer array nums and an integer n, add/patch element
//s to the array such that any number in range [1, n] inclusive can be formed by t
//he sum of some elements in the array. Return the minimum number of patches requi
//red.
//
// Example 1:
//
//
//Input: nums = [1,3], n = 6
//Output: 1
//Explanation:
//Combinations of nums are [1], [3], [1,3], which form possible sums of: 1, 3, 4
//.
//Now if we add/patch 2 to nums, the combinations are: [1], [2], [3], [1,3], [2,
//3], [1,2,3].
//Possible sums are 1, 2, 3, 4, 5, 6, which now covers the range [1, 6].
//So we only need 1 patch.
//
// Example 2:
//
//
//Input: nums = [1,5,10], n = 20
//Output: 2
//Explanation: The two patches can be [2, 4].
//
//
// Example 3:
//
//
//Input: nums = [1,2,2], n = 5
//Output: 0
// Related Topics Greedy
// 👍 595 👎 84

// 2021-02-18 18:45:55

//leetcode submit region begin(Prohibit modification and deletion)
func minPatches(nums []int, n int) int {
	boundary, l := 1, len(nums)
	res, i := 0, 0

	for boundary <= n {
		if i < l && boundary >= nums[i] {
			boundary += nums[i]
			i++
		} else {
			boundary += boundary
			res++
		}

	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
