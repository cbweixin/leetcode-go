package main

import "strings"

// 实现一个基本的计算器来计算简单的表达式字符串。
//
// 表达式字符串只包含非负整数，算符 +、-、*、/ ，左括号 ( 和右括号 ) 。整数除法需要 向下截断 。
//
// 你可以假定给定的表达式总是有效的。所有的中间结果的范围为 [-2³¹, 2³¹ - 1] 。
//
//
//
// 示例 1：
//
//
// 输入：s = "1+1"
// 输出：2
//
//
// 示例 2：
//
//
// 输入：s = "6-4/2"
// 输出：4
//
//
// 示例 3：
//
//
// 输入：s = "2*(5+5*2)/3+(6/2+8)"
// 输出：21
//
//
// 示例 4：
//
//
// 输入：s = "(2+6*3+5-(3*14/7+2)*5)+3"
// 输出：-12
//
//
// 示例 5：
//
//
// 输入：s = "0"
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= s <= 10⁴
// s 由整数、'+'、'-'、'*'、'/'、'(' 和 ')' 组成
// s 是一个 有效的 表达式
//
//
//
//
// 进阶：你可以在不使用内置库函数的情况下解决此问题吗？
//
// Related Topics 栈 递归 数学 字符串 👍 128 👎 0

// 2022-08-01 12:02:59
// leetcode submit region begin(Prohibit modification and deletion)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	operands, operators := make([]int, 0), make([]int, 0)

	compute := func(b, a int, op int) int {
		if op == '+' {
			return a + b
		} else if op == '-' {
			return a - b
		} else if op == '*' {
			return a * b
		}
		return a / b
	}

	pop := func(arr *[]int) int {
		m := len(*arr) - 1
		t := (*arr)[m]
		*arr = (*arr)[:m]
		return t
	}

	getPriority := func(op int) int {
		if op == '(' {
			return 0
		} else if op == '+' || op == '-' {
			return 1
		}
		return 2
	}

	temp, l := 0, len(s)-1

	for i, c := range s {
		if c <= '9' && c >= '0' {
			temp = temp*10 + int(c-'0')
			if i == l {
				operands = append(operands, temp)
				temp = 0
			}
		} else {
			operands = append(operands, temp)
			temp = 0
			if c == '(' {
				operators = append(operators, int(c))
			} else if c == ')' {
				for operators[len(operators)-1] != '(' {
					t := compute(pop(&operands), pop(&operands), pop(&operators))
					operands = append(operands, t)
				}
				// pop the last (
				pop(&operators)
			} else {
				// if c is '+-*/'
				for len(operators) > 0 && getPriority(int(c)) <= getPriority(operators[len(operators)-1]) {
					t := compute(pop(&operands), pop(&operands), pop(&operators))
					operands = append(operands, t)
				}
				operators = append(operators, int(c))
			}
		}
	}

	for len(operators) > 0 {
		t := compute(pop(&operands), pop(&operands), pop(&operators))
		operands = append(operands, t)
	}

	if len(operands) == 0 {
		return 0
	}
	return pop(&operands)
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "1 + (2-3)*(5-6/2)"
	println(calculate(s))
}
