package main

import "math"

//Given a single positive integer x, we will write an expression of the form x (
//op1) x (op2) x (op3) x ... where each operator op1, op2, etc. is either addition
//, subtraction, multiplication, or division (+, -, *, or /). For example, with x 
//= 3, we might write 3 * 3 / 3 + 3 - 3 which is a value of 3. 
//
// When writing such an expression, we adhere to the following conventions: 
//
// 
// The division operator (/) returns rational numbers. 
// There are no parentheses placed anywhere. 
// We use the usual order of operations: multiplication and division happens bef
//ore addition and subtraction. 
// It's not allowed to use the unary negation operator (-). For example, "x - x"
// is a valid expression as it only uses subtraction, but "-x + x" is not because 
//it uses negation. 
// 
//
// We would like to write an expression with the least number of operators such 
//that the expression equals the given target. Return the least number of operator
//s used. 
//
// 
//
// 
// Example 1: 
//
// 
//Input: x = 3, target = 19
//Output: 5
//Explanation: 3 * 3 + 3 * 3 + 3 / 3.  The expression contains 5 operations.
// 
//
// Example 2: 
//
// 
// 
//Input: x = 5, target = 501
//Output: 8
//Explanation: 5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5.  The expression contains 8 ope
//rations.
// 
//
// 
// Example 3: 
//
// 
//Input: x = 100, target = 100000000
//Output: 3
//Explanation: 100 * 100 * 100 * 100.  The expression contains 3 operations. 
//
// 
// 
// 
// 
//
// Note: 
//
// 
// 2 <= x <= 100 
// 1 <= target <= 2 * 10^8 
// 
//
// 
// 
// 
// 
// 
// Related Topics Math Dynamic Programming 
// 👍 190 👎 50

// 2020-10-30 11:47:10

//leetcode submit region begin(Prohibit modification and deletion)
func leastOpsExpressTarget(x int, target int) int {
	memo := make(map[int]int)

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var dfs func(int) int

	dfs = func(goal int) int {
		if v, ok := memo[goal]; ok {
			return v
		}

		if goal == x {
			memo[goal] = 0
			return 0
		}

		if x > goal{
			memo[goal] = min(goal * 2 - 1, (x-goal)*2)
			return memo[goal]
		}

		sums, cnt := x, 0
		for sums < goal {
			sums *= goal
			cnt ++
		}

		if sums == goal {
			memo[sums] = cnt
			return cnt
		}

		l, r := math.MaxInt32, math.MaxInt32
		if sums - goal < goal {
			l = dfs(sums-goal) + cnt
		}
		r = dfs(int(sums/goal)) + cnt +1

		memo[goal] = min(l,r) + 1
		return memo[goal]
	}

	return dfs(target)
}

//leetcode submit region end(Prohibit modification and deletion)
