package main

import "fmt"

//给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
//
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
//
//
//
// 示例 1：
//
//
//输入：n = 13
//输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
//
//
// 示例 2：
//
//
//输入：n = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 5 * 10⁴
//
//
// Related Topics 深度优先搜索 字典树 👍 447 👎 0

//2023-05-06 13:31:06

// leetcode submit region begin(Prohibit modification and deletion)
func lexicalOrder2(n int) []int {
	res, cur := make([]int, n), 1

	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur++
		}
	}
	return res
}

func lexicalOrder(n int) []int {

	var dfs func(int, int, *[]int)

	dfs = func(cur int, limit int, res *[]int) {
		if cur > limit {
			return
		}
		*res = append(*res, cur)
		for i := 0; i < 10; i++ {
			dfs(cur*10+i, limit, res)
		}
	}

	res := make([]int, 0)
	for i := 1; i < 10; i++ {
		dfs(i, n, &res)
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(lexicalOrder(23))
}
