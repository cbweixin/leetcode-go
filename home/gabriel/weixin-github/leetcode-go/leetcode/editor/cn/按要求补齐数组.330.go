package main

/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
func minPatches(nums []int, n int) int {
	boundary, cnt, i := 1, 0, 0

	for boundary < n+1 {
		if nums[i] <= boundary {
			boundary += nums[i]
			i++
		} else {
			boundary += boundary
			cnt++

		}

	}
	return cnt

}

// @lc code=end
