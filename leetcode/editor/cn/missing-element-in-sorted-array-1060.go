package main

func missingElement1(nums []int, k int) int {
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

func missingElement(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid]-nums[0]-mid >= k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[0] + k + r
}
