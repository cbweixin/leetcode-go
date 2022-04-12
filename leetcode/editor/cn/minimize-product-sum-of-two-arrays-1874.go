package main

import "sort"

// sort
func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res, l := 0, len(nums1)

	for i := 0; i < l; i++ {
		res += nums1[i] * nums2[l-i-1]
	}

	return res
}
