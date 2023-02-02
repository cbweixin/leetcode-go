package main

//给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。
//
//
//
// 示例 1：
//
//
//输入：s = "abab"
//输出："bab"
//解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是
//"bab"。
//
//
// 示例 2：
//
//
//输入：s = "leetcode"
//输出："tcode"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 4 * 10⁵
// s 仅含有小写英文字符。
//
//
// Related Topics 双指针 字符串 👍 74 👎 0

//2023-01-31 21:31:46

// leetcode submit region begin(Prohibit modification and deletion)
func lastSubstring(s string) string {
	left, right, step := 0, 1, 0
	for right+step < len(s) {
		if s[left+step] < s[right+step] {
			left, right, step = right, right+1, 0
		} else if s[left+step] > s[right+step] {
			right, step = right+step+1, 0
		} else {
			step += 1
		}
	}
	return s[left:]

}

//leetcode submit region end(Prohibit modification and deletion)
