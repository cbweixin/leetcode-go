package main

import "fmt"

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 3038 👎 0

//2023-01-30 14:04:10

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	res, path := make([]string, 0), make([]byte, 0)
	var gen func(int, int, []byte)

	gen = func(left, right int, path []byte) {
		if left == n && right == n {
			res = append(res, string(path))
			return
		}
		if left < n {
			gen(left+1, right, append(path, '('))
		}
		if right < left {
			gen(left, right+1, append(path, ')'))
		}
	}
	gen(0, 0, path)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(generateParenthesis(3))
}
