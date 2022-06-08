package main

//给出一个字符串 s 和一个整数 k，若这个字符串是不是一个「k 回文 」，则返回 true 。
//
// 如果可以通过从字符串中删去最多 k 个字符将其转换为回文，那么这个字符串就是一个「k 回文 」。
//
//
//
// 示例 1：
//
//
//输入：s = "abcdeca", k = 2
//输出：true
//解释：删去字符 “b” 和 “e”。
//
//
// 示例 2:
//
//
//输入：s = "abbababa", k = 1
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 中只含有小写英文字母
// 1 <= k <= s.length
//
// Related Topics 字符串 动态规划 👍 42 👎 0

// 2022-06-08 07:02:41
//leetcode submit region begin(Prohibit modification and deletion)
func isValidPalindrome(s string, k int) bool {
	l := len(s)
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}
	t := reverse(s)

	dp := make([][]int, l+1)
	for x := range dp {
		dp[x] = make([]int, l+1)
		for i := 1; i < l+1; i++ {
			dp[x][i] = 0
		}
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for i := 1; i < l+1; i++ {
		for j := 1; j < l+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return l-dp[l][l] <= k
}

//leetcode submit region end(Prohibit modification and deletion)
