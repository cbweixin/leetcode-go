package main

import (
	"fmt"
	"math"
)

func maximizeSweetness(sweetness []int, k int) int {
	l, r := math.MaxInt, 0
	for _, v := range sweetness {
		if l > v {
			l = v
		}
		r += v

	}

	splitParts := func(mid int) int {
		s, parts := 0, 0
		for _, v := range sweetness {
			s += v
			if s > mid {
				s = 0
				parts++
			}
		}
		return parts
	}

	for l < r {
		mid := (l + r) >> 1
		if splitParts(mid) >= k+1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	sweetness := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(maximizeSweetness(sweetness, 5))
}
