package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	partition := func(left, right, pivot int) int {
		nums[pivot], nums[right] = nums[right], nums[pivot]
		k := left

		for i := left; i < right; i++ {
			if nums[i] > nums[right] {
				nums[i], nums[k] = nums[k], nums[i]
				k++
			}
		}

		nums[k], nums[right] = nums[right], nums[k]
		return k
	}

	left, right := 0, len(nums)-1

	for left <= right {
		rnk := partition(left, right, left+rand.Intn(right-left+1))
		if rnk == k-1 {
			return nums[rnk]
		}
		if rnk > k-1 {
			right = rnk - 1
		} else {
			left = rnk + 1
		}
	}

	return -1

}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}
