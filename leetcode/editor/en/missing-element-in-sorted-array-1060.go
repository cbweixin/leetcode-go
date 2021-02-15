package main

import "fmt"

func missingElement(nums []int, k int) int {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid]-nums[0]-mid < k {
			left = mid
		} else {
			right = mid - 1
		}

	}

	return nums[0] + k + left

}

func main() {
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 1))

}
