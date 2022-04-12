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

func minProductSum1(nums1 []int, nums2 []int) int {
	// 1 <= nums1[i], nums2[i] <= 100
	bucket1, bucket2 := make([]int, 101), make([]int, 102)
	l := len(nums1)

	for i := 0; i < l; i++ {
		bucket1[nums1[i]] += 1
		bucket2[nums2[i]] += 1
	}

	left, right, res := 0, 100, 0
	for left < 101 && right >= 0 {
		if bucket1[left] > 0 && bucket2[right] > 0 {
			res += left * right
			bucket1[left]--
			bucket2[right]--
		} else if bucket1[left] > 0 {
			right--
		} else if bucket2[right] > 0 {
			left++
		} else {
			left++
			right--
		}
	}
	return res
}
