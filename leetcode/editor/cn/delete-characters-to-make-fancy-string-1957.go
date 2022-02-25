package main

import (
	"fmt"
	"strings"
)

//一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。
//
// 给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。
//
// 请你返回删除后的字符串。题目数据保证答案总是 唯一的 。
//
//
//
// 示例 1：
//
//
//输入：s = "leeetcode"
//输出："leetcode"
//解释：
//从第一组 'e' 里面删除一个 'e' ，得到 "leetcode" 。
//没有连续三个相同字符，所以返回 "leetcode" 。
//
//
// 示例 2：
//
//
//输入：s = "aaabaaaa"
//输出："aabaa"
//解释：
//从第一组 'a' 里面删除一个 'a' ，得到 "aabaaaa" 。
//从第二组 'a' 里面删除两个 'a' ，得到 "aabaa" 。
//没有连续三个相同字符，所以返回 "aabaa" 。
//
//
// 示例 3：
//
//
//输入：s = "aab"
//输出："aab"
//解释：没有连续三个相同字符，所以返回 "aab" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 只包含小写英文字母。
//
// Related Topics 字符串 👍 14 👎 0

// 2022-02-25 09:06:40
//leetcode submit region begin(Prohibit modification and deletion)
func makeFancyString_tle(s string) string {
	var b strings.Builder
	for _, c := range s {
		chars := []rune(b.String())
		l := len(chars)
		if l >= 2 && chars[l-1] == c && chars[l-2] == c {
			continue
		}
		b.WriteString(string(c))
	}

	return b.String()
}

func makeFancyString2(s string) string {
	res := make([]byte, 0, len(s))
	cur := 0
	curChar := s[0]
	for i := 0; i < len(s); i++ {
		if curChar != s[i] {
			curChar = s[i]
			cur = 0
		}
		cur++
		if cur <= 2 {
			res = append(res, curChar)
		}
	}

	return string(res)

}

func makeFancyString(s string) string {
	// create a byte array instead of a string
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		l := len(res)
		// s[i] is uint8, res[l] is byte (alias for uint8)
		if l >= 2 && res[l-1] == s[i] && res[l-2] == s[i] {
			continue
		}
		res = append(res, s[i])
	}

	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "leeetcode"
	fmt.Println(makeFancyString(s))
}
