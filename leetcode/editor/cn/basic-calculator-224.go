package main

import "strings"

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
//
//
// 示例 1：
//
//
// 输入：s = "1 + 1"
// 输出：2
//
//
// 示例 2：
//
//
// 输入：s = " 2-1 + 2 "
// 输出：3
//
//
// 示例 3：
//
//
// 输入：s = "(1+(4+5+2)-3)+(6+8)"
// 输出：23
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
// '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
// '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
// 输入中不存在两个连续的操作符
// 每个数字和运行的计算将适合于一个有符号的 32位 整数
//
//
// Related Topics 栈 递归 数学 字符串 👍 802 👎 0

// 2022-08-01 09:59:08
// leetcode submit region begin(Prohibit modification and deletion)
func calculate1(s string) int {
	var b strings.Builder
	b.WriteString(s)
	b.WriteString("+")
	sign, temp, res := 1, 0, 0
	ops := []int{1}
	for _, c := range b.String() {
		if c == '(' {
			ops = append(ops, sign)
		} else if c == ')' {
			ops = ops[:len(ops)-1]
		} else if c == '+' || c == '-' {
			res += sign * temp
			if c == '-' {
				sign = -1 * ops[len(ops)-1]
			} else {
				sign = ops[len(ops)-1]
			}
			temp = 0
		} else if c >= '0' && c <= '9' {
			temp = temp*10 + int(c-'0')
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "1 -(1-2) + 3 -(4-5-5)"
	println(calculate1(s))
}
