package main

import (
	"container/heap"
	"fmt"
)

// 给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
//
//
//
// 示例 1:
//
//
//
//
// 输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
// 输出: 4
// 解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
//
//
// 示例 2:
//
//
//
//
// 输入: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
// 输出: 10
//
//
//
//
// 提示:
//
//
// m == heightMap.length
// n == heightMap[i].length
// 1 <= m, n <= 200
// 0 <= heightMap[i][j] <= 2 * 10⁴
//
//
//
//
// Related Topics 广度优先搜索 数组 矩阵 堆（优先队列） 👍 648 👎 0

// 2023-01-16 22:27:59
// leetcode submit region begin(Prohibit modification and deletion)
func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m == 0 {
		return 0
	}
	n := len(heightMap[0])
	if n == 0 {
		return 0
	}
	pq := make(ConeHeap, 0)
	heap.Init(&pq)
	isVisited := make([][]bool, m)
	for i := range isVisited {
		isVisited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				pq.Push(cone{height: heightMap[i][j], x: i, y: j})
				isVisited[i][j] = true
			}
		}
	}

	trap := 0
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for pq.Len() > 0 {
		topCone := pq.Pop().(cone)
		for _, v := range [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			nX, nY := topCone.x+v[0], topCone.y+v[1]
			if nX >= 0 && nX < m && nY >= 0 && nY < n && !isVisited[nX][nY] {
				trap += max(0, topCone.height-heightMap[nX][nY])
				pq.Push(
					cone{
						height: max(topCone.height, heightMap[nX][nY]),
						x:      nX,
						y:      nY,
					},
				)
				isVisited[nX][nY] = true
			}
		}

	}

	return trap
}

type cone struct {
	height int
	x      int
	y      int
}

// knowledge how to implement heap
type ConeHeap []cone

func (h ConeHeap) Len() int           { return len(h) }
func (h ConeHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h ConeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ConeHeap) Push(x interface{}) {
	*h = append(*h, x.(cone))
}
func (h *ConeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(trapRainWater([][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}))

}
