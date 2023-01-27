package main

import "strings"

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 8635 👎 0

//2023-01-27 12:05:03

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	pos, res := make(map[uint8]int), 0
	var b strings.Builder
	b.WriteString(s)
	b.WriteString("#")
	s = b.String()
	l := len(s)

	left, right := 0, 0
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for right < l {
		res = max(res, right-left)
		v, ok := pos[s[right]]
		if ok && v >= left {
			left = v + 1
		}
		pos[s[right]] = right
		right += 1

	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
