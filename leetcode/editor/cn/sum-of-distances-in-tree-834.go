package main

// 给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。
//
// 第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。
//
// 返回一个表示节点 i 与其他所有节点距离之和的列表 ans。
//
// 示例 1:
//
//
// 输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
// 输出: [8,12,6,10,10,10]
// 解释:
// 如下为给定的树的示意图：
//  0
// / \
// 1   2
//   /|\
//  3 4 5
//
// 我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
// 也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
//
//
// 说明: 1 <= N <= 10000
// Related Topics 树 深度优先搜索 图 动态规划 👍 325 👎 0

// 2022-01-29 12:56:21
// leetcode submit region begin(Prohibit modification and deletion)
func sumOfDistancesInTree(n int, edges [][]int) []int {
	distSum := make([]int, n)
	nodeSum := make([]int, n)
	for i := 0; i < n; i++ {
		nodeSum[i] = 1
	}
	graph := make([][]int, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	var postOrder func(int, int)
	postOrder = func(node, parent int) {
		for _, nei := range graph[node] {
			if nei == parent {
				continue
			}
			postOrder(nei, node)
			nodeSum[node] += nodeSum[nei]
			distSum[node] += nodeSum[nei] + distSum[nei]
		}
	}

	var preOrder func(int, int)
	preOrder = func(node, parent int) {
		for _, nei := range graph[node] {
			if nei == parent {
				continue
			}
			distSum[nei] = distSum[node] + n - (nodeSum[nei] << 1)
			preOrder(nei, node)
		}
	}
	postOrder(0, -1)
	preOrder(0, -1)

	return distSum
}

// leetcode submit region end(Prohibit modification and deletion)
