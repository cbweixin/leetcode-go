package main

import "fmt"

//Consider the string s to be the infinite wraparound string of "abcdefghijklmno
//pqrstuvwxyz", so s will look like this: "...zabcdefghijklmnopqrstuvwxyzabcdefghi
//jklmnopqrstuvwxyzabcd....".
//
// Now we have another string p. Your job is to find out how many unique non-emp
//ty substrings of p are present in s. In particular, your input is the string p a
//nd you need to output the number of different non-empty substrings of p in the s
//tring s.
//
// Note: p consists of only lowercase English letters and the size of p might be
// over 10000.
//
// Example 1:
//
//Input: "a"
//Output: 1
//
//Explanation: Only the substring "a" of string "a" is in the string s.
//
//
//
// Example 2:
//
//Input: "cac"
//Output: 2
//Explanation: There are two substrings "a", "c" of string "cac" in the string s
//.
//
//
//
// Example 3:
//
//Input: "zab"
//Output: 6
//Explanation: There are six substrings "z", "a", "b", "za", "ab", "zab" of stri
//ng "zab" in the string s.
//
// Related Topics Dynamic Programming
// 👍 724 👎 97

// 2021-01-21 09:38:08

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstringInWraproundString_wrong(p string) int {
	L := len(p)
	dp := make([]int, L)
	for i := range dp {
		dp[i] = 1
	}
	l := 1

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i := 1; i < L; i++ {
		if (int(p[i])-int('a')+1)%26 == int(p[i-1])-int('a') {
			l++
		} else {
			l = 1
		}
		// this is wrong, because for "cac", 'c' would be calcuated twice
		dp[i] = max(dp[i], l)
	}

	res := 0
	for _, v := range dp {
		res += v
	}

	return res

}

func findSubstringInWraproundString(p string) int {
	L := len(p)
	dp := make(map[byte]int)
	for i := 0; i < L; i++ {
		dp[p[i]] = 1

	}
	l := 1

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	getOrder := func(i int) int {
		return int(p[i]) - int('a')
	}
	for i := 1; i < L; i++ {
		if (getOrder(i-1)+1)%26 == getOrder(i) {
			l++
		} else {
			l = 1
		}
		// this is wrong, because for "cac", 'c' would be calcuated twice
		dp[p[i]] = max(dp[p[i]], l)
	}

	res := 0
	for _, v := range dp {
		res += v
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findSubstringInWraproundString("cac"))
}
