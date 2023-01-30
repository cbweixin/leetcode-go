package main

//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
//在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
//
// 示例 1:
//
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//
//
// 示例 2:
//
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//
//
// 示例 3:
//
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//
//
// 提示：
//
//
// 1 <= temperatures.length <= 10⁵
// 30 <= temperatures[i] <= 100
//
//
// Related Topics 栈 数组 单调栈 👍 1384 👎 0

//2023-01-29 21:22:14

// leetcode submit region begin(Prohibit modification and deletion)
// stack, mono-stack
func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	res, stack := make([]int, l), make([]int, 0)

	for i, v := range temperatures {
		t := stack[len(stack)-1]
		for len(stack) > 0 && v > temperatures[t] {
			res[t] = i - t
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
