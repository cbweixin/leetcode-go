package main

//给你这棵「无向树」，请你测算并返回它的「直径」：这棵树上最长简单路径的 边数。
//
// 我们用一个由所有「边」组成的数组 edges 来表示一棵无向树，其中 edges[i] = [u, v] 表示节点 u 和 v 之间的双向边。
//
// 树上的节点都已经用 {0, 1, ..., edges.length} 中的数做了标记，每个节点上的标记都是独一无二的。
//
//
//
// 示例 1：
//
//
//
// 输入：edges = [[0,1],[0,2]]
//输出：2
//解释：
//这棵树上最长的路径是 1 - 0 - 2，边数为 2。
//
//
// 示例 2：
//
//
//
// 输入：edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
//输出：4
//解释：
//这棵树上最长的路径是 3 - 2 - 1 - 4 - 5，边数为 4。
//
//
//
//
// 提示：
//
//
// 0 <= edges.length < 10^4
// edges[i][0] != edges[i][1]
// 0 <= edges[i][j] <= edges.length
// edges 会形成一棵无向树
//
// Related Topics 树 深度优先搜索 广度优先搜索 图 拓扑排序 👍 91 👎 0

// 2022-06-14 06:44:37
//leetcode submit region begin(Prohibit modification and deletion)

type MTreeNode struct {
	Denomin  int
	Children []MTreeNode
}

func treeDiameter(edges [][]int) int {
	l := len(edges)
	nodes := make([]MTreeNode, l+1)
	for i := 0; i < l; i++ {
		nodes[i] = MTreeNode{
			Denomin:  i,
			Children: make([]MTreeNode, 0),
		}
	}
	for _, v := range edges {
		nodes[v[0]].Children = append(nodes[v[0]].Children, nodes[v[1]])
		nodes[v[1]].Children = append(nodes[v[1]].Children, nodes[v[0]])
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	res := 0
	var dfs func(int, int) int
	dfs = func(n, pre int) int {

		d1, d2 := 0, 0
		for _, node := range nodes[n].Children {
			if node.Denomin == pre {
				continue
			}
			d := dfs(node.Denomin, n)
			if d > d1 {
				d1, d2 = d, d1
			} else if d > d2 {
				d2 = d
			}
		}

		res = max(res, d1+d2)
		return d1 + 1
	}

	dfs(0, -1)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	edges := [][]int{{0, 1}, {0, 2}}
	println(treeDiameter(edges))
}
