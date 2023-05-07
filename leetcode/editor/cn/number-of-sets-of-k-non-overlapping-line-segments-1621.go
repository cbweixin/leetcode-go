package main

import "fmt"

//给你一维空间的 n 个点，其中第 i 个点（编号从 0 到 n-1）位于 x = i 处，请你找到 恰好 k 个不重叠 线段且每个线段至少覆盖两个点的方案数
//。线段的两个端点必须都是 整数坐标 。这 k 个线段不需要全部覆盖全部 n 个点，且它们的端点 可以 重合。
//
// 请你返回 k 个不重叠线段的方案数。由于答案可能很大，请将结果对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：5
//解释：
//如图所示，两个线段分别用红色和蓝色标出。
//上图展示了 5 种不同的方案 {(0,2),(2,3)}，{(0,1),(1,3)}，{(0,1),(2,3)}，{(1,2),(2,3)}，{(0,1),
//(1,2)} 。
//
// 示例 2：
//
//
//输入：n = 3, k = 1
//输出：3
//解释：总共有 3 种不同的方案 {(0,1)}, {(0,2)}, {(1,2)} 。
//
//
// 示例 3：
//
//
//输入：n = 30, k = 7
//输出：796297179
//解释：画 7 条线段的总方案数为 3796297200 种。将这个数对 10⁹ + 7 取余得到 796297179 。
//
//
// 示例 4：
//
//
//输入：n = 5, k = 3
//输出：7
//
//
// 示例 5：
//
//
//输入：n = 3, k = 2
//输出：1
//
//
//
// 提示：
//
//
// 2 <= n <= 1000
// 1 <= k <= n-1
//
//
// Related Topics 数学 动态规划 👍 54 👎 0

//2023-05-06 15:51:34

func get_combin(n int, m int) int {
	cache := make([][]int, n+1)
	mod := 1000000007

	for i := range cache {
		cache[i] = make([]int, m+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var combine func(int, int) int

	// c(n,k) = c(n-1,k) + c(n-1, k-1)
	combine = func(n int, k int) int {
		if k == 0 {
			return 1
		} else if n == 0 {
			return 0
		} else if cache[n][k] != -1 {
			return cache[n][k]
		} else {
			cache[n][k] = (combine(n-1, k-1) + combine(n-1, k)) % mod
		}

		return cache[n][k]
	}

	t := combine(n, m)
	//return combine(n, m)
	//fmt.Println("%v", cache)
	return t

}

func numberOfSets(n int, k int) int {
	return get_combin(n+k-1, k<<1)
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numberOfSets(4, 2))
}
