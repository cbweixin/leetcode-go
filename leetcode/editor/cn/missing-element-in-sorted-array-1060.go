package main

func missingElement(nums []int, k int) int {
	res, cur, i, l := -1, nums[0]+1, 1, len(nums)
	for k > 0 {
		if (i < l && cur != nums[i]) || i >= l {
			res = cur
			k--
		} else {
			i++
		}
		cur++
	}
	return res
}
