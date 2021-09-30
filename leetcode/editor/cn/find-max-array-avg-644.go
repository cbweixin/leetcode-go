package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	min, max := float64(math.MaxInt32), float64(math.MinInt32)
	for _, v := range nums {
		k := float64(v)
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}

	hasAvg := func(avg float64) bool {
		sum := 0.0
		for i := 0; i < k; i++ {
			sum += float64(nums[i]) - avg
		}
		if sum >= 0 {
			return true
		}

		prev := 0.0
		for i := k; i < len(nums); i++ {
			sum += float64(nums[i]) - avg
			prev += float64(nums[i-k]) - avg

			if prev < 0 {
				sum -= prev
				prev = 0
			}
			if sum >= 0 {
				return true
			}

		}
		return false
	}

	delta := 1.0
	for delta > 0.000005 {
		mid := (max + min) * 0.5
		if hasAvg(mid) {
			min = mid
		} else {
			max = mid
		}

		delta = max - min

	}

	return min

}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	r := findMaxAverage(nums, k)
	fmt.Print(r)

}
