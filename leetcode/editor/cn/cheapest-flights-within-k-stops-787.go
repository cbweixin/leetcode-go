package main

import (
	"fmt"
	"math"
)

//有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
//市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
//宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
//
//
// 示例 1：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//输出: 200
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
//
// 示例 2：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 0
//输出: 500
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 动态规划 最短路 堆（优先队列） 👍 574 👎 0

//2023-05-05 16:04:53

// leetcode submit region begin(Prohibit modification and deletion)
func findCheapestPrice_bfs(n int, flights [][]int, src int, dst int, k int) int {
	costs := make([][]int, n)
	curMin := make([][]int, k+1)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt
		}
	}

	for i := range curMin {
		curMin[i] = make([]int, n)
	}

	for _, f := range flights {
		costs[f[0]][f[1]] = f[2]
	}

	minCost := math.MaxInt
	//var null struct{}
	//seen := make(map[int]struct{})
	//seen[src] = null

	que := make([]int, 0)
	que = append(que, src)

	curMin[0][src] = 0
	step := 0

	for len(que) > 0 && step <= k {
		m := len(que)
		for i := 0; i < m; i++ {
			s := que[0]
			que = que[1:]
			for d := 0; d < n; d++ {
				if costs[s][d] == math.MaxInt {
					continue
				}
				if step == 0 {
					curMin[step][d] = costs[s][d]
				} else {
					t := curMin[step-1][s] + costs[s][d]
					if t < curMin[step][d] || curMin[step][d] == 0 {
						curMin[step][d] = t
					} else {
						continue
					}
				}
				if curMin[step][d] > minCost {
					continue
				}

				if d == dst {
					minCost = curMin[step][d]
				}
				que = append(que, d)
			}
		}
		step++
	}
	if minCost == math.MaxInt {
		return -1
	}

	return minCost

}
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp, temp := make([]int, n), make([]int, n)

	for i := range dp {
		dp[i] = 100000
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	reset_temp := func() {
		for i := range temp {
			temp[i] = 100000
		}
		temp[src] = 0
	}
	dp[src] = 0
	for i := 0; i < k+1; i++ {
		reset_temp()
		for _, f := range flights {
			temp[f[1]] = min(temp[f[1]], dp[f[0]]+f[2])
		}
		for i := range temp {
			dp[i] = temp[i]
		}

	}
	//fmt.Printf("%v", dp)
	if dp[dst] == 100000 {
		return -1
	}
	return dp[dst]
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	n := 5
	flight := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}}
	src := 0
	dst := 4
	k := 2
	fmt.Println(findCheapestPrice(n, flight, src, dst, k))
}
