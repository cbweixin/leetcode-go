package main

// 给定一个数组 trees，其中 trees[i] = [xi, yi] 表示树在花园中的位置。
//
// 你被要求用最短长度的绳子把整个花园围起来，因为绳子很贵。只有把 所有的树都围起来，花园才围得很好。
//
// 返回恰好位于围栏周边的树木的坐标。
//
// 示例 1:
//
//
//
//
// 输入: points = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
// 输出: [[1,1],[2,0],[3,3],[2,4],[4,2]]
//
// 示例 2:
//
//
//
//
// 输入: points = [[1,2],[2,2],[4,2]]
// 输出: [[4,2],[2,2],[1,2]]
//
//
//
// 注意:
//
//
// 1 <= points.length <= 3000
// points[i].length == 2
// 0 <= xi, yi <= 100
// 所有给定的点都是 唯一 的。
//
//
// Related Topics 几何 数组 数学 👍 210 👎 0

// 2022-12-31 12:34:40
// leetcode submit region begin(Prohibit modification and deletion)
func outerTrees(trees [][]int) [][]int {
	cross := func(p, q, r []int) int {
		return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
	}

	visited := make([]bool, len(trees))
	leftMost := 0
	for k, v := range trees {
		if v[0] < trees[leftMost][0] {
			leftMost = k
		}
	}
	p := leftMost
	var res [][]int
	for {
		q := (p + 1) % len(trees)
		for r, tree := range trees {
			// if r is on the left side of p,q
			if cross(trees[p], trees[q], tree) < 0 {
				q = r
			}
		}
		// if there exists i, which is at the line formed by p, q
		for i, b := range visited {
			if !b && i != p && i != q && cross(trees[p], trees[q], trees[i]) == 0 {
				res = append(res, trees[i])
				visited[i] = true
			}
		}

		if !visited[q] {
			res = append(res, trees[q])
			visited[q] = true
		}

		p = q
		if p == leftMost {
			break
		}

	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
