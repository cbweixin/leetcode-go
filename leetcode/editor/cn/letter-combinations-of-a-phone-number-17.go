package main

import (
	"fmt"
	"strings"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 2283 👎 0

//2023-01-30 09:59:54

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	var lookup = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := make([]string, 0)
	var dfs func(int, string)

	dfs = func(idx int, s string) {
		if idx == len(digits) {
			res = append(res, s)
			return
		}
		for _, v := range lookup[digits[idx]-'0'] {
			var b strings.Builder
			b.WriteString(s)
			b.WriteRune(v)
			dfs(idx+1, b.String())

		}
		return

	}

	if len(digits) == 0 {
		return res
	}
	dfs(0, "")
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(letterCombinations("23"))
}
