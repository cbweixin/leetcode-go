package main

import "math"

func assignBikes2(workers [][]int, bikes [][]int) int {
	l, m := len(workers), len(bikes)
	upper := 1 << m
	dp := make([][]int, l+1)

	for i := 0; i < l+1; i++ {
		dp[i] = make([]int, upper)
		for j := 0; j < upper; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0

	count_bits := func(mask int) int {
		res := 0
		for mask > 0 {
			res++
			mask &= mask - 1
		}

		return res
	}

	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	abs := func(x int) int {
		if x <= 0 {
			return -x
		}
		return x
	}

	dist := func(w, b []int) int {
		return abs(w[0]-b[0]) + abs(w[1]-b[1])
	}

	res := math.MaxInt
	for i := 1; i < l+1; i++ {
		for mask := 0; mask < upper; mask++ {
			if count_bits(mask) != i {
				continue
			}
			for j := 0; j < m; j++ {
				if mask&(1<<j) > 0 {
					dp[i][mask] = min(dp[i][mask], dp[i-1][mask^(1<<j)]+dist(workers[i-1], bikes[j]))
					if i == l {
						res = min(res, dp[i][mask])
					}
				}
			}

		}
	}
	return res
}
func main() {
	workers := [][]int{{0, 0}, {2, 1}}
	bikes := [][]int{{1, 2}, {3, 3}}
	println(assignBikes(workers, bikes))

}
