package main

//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 6079 👎 0

//2023-01-27 15:44:53

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	l := len(s)
	palindromeLen := func(left int, right int) string {
		res := string(s[left])
		for left >= 0 && right < l && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left, right = left-1, right+1
		}
		return res
	}

	res := string(s[0])
	max := func(x, y string) string {
		if len(x) >= len(y) {
			return x
		}
		return y
	}

	for i := 0; i < l-1; i++ {
		a, b := palindromeLen(i, i), palindromeLen(i, i+1)
		t := max(a, b)
		if len(t) > len(res) {
			res = t
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
