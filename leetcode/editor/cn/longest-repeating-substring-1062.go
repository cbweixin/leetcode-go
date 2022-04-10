package main

import "fmt"

func longestRepeatingSubstring(s string) int {
	var l = len(s)
	var dp [][]int
	for _, arr := range dp {
		for j := range arr {
			arr[j] = 0
		}
	}
	res := 0
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i := 1; i < l+1; i++ {
		for j := i + 1; j < l+1; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}
func main() {
	s := "abcd"
	fmt.Println(longestRepeatingSubstring(s))
}
