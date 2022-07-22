package main

import (
	"math"
	"sort"
)

// 给定两个正整数数组 boxes 和 warehouse ，分别包含单位宽度的箱子的高度，以及仓库中 n 个房间各自的高度。仓库的房间分别从 0 到 n -
// 1 自左向右编号， warehouse[i] （索引从 0 开始）是第 i 个房间的高度。
//
// 箱子放进仓库时遵循下列规则：
//
//
// 箱子不可叠放。
// 你可以重新调整箱子的顺序。
// 箱子只能从左向右推进仓库中。
// 如果仓库中某房间的高度小于某箱子的高度，则这个箱子和之后的箱子都会停在这个房间的前面。
//
//
// 你最多可以在仓库中放进多少个箱子？
//
//
//
// 示例 1：
//
//
//
//
// 输入：boxes = [4,3,4,1], warehouse = [5,3,3,4,1]
// 输出：3
// 解释：
//
// 我们可以先把高度为 1 的箱子放入 4 号房间，然后再把高度为 3 的箱子放入 1 号、 2 号或 3 号房间，最后再把高度为 4 的箱子放入 0 号房间。
//
// 我们不可能把所有 4 个箱子全部放进仓库里。
//
// 示例 2：
//
//
//
//
// 输入：boxes = [1,2,2,3,4], warehouse = [3,4,1,2]
// 输出：3
// 解释：
//
// 我们注意到，不可能把高度为 4 的箱子放入仓库中，因为它不能通过高度为 3 的房间。
// 而且，对于最后两个房间 2 号和 3 号来说，只有高度为 1 的箱子可以放进去。
// 我们最多可以放进 3 个箱子，如上图所示。黄色的箱子也可以放入 2 号房间。
// 交换橙色和绿色箱子的位置，或是将这两个箱子与红色箱子交换位置，也是可以的。
//
// 示例 3：
//
//
// 输入：boxes = [1,2,3], warehouse = [1,2,3,4]
// 输出：1
// 解释：由于第一个房间的高度为 1，我们只能放进高度为 1 的箱子。
//
//
//
//
// 提示：
//
//
// n == warehouse.length
// 1 <= boxes.length, warehouse.length <= 10^5
// 1 <= boxes[i], warehouse[i] <= 10^9
//
// Related Topics 贪心 数组 排序 👍 13 👎 0

// 2022-07-22 08:53:15
// leetcode submit region begin(Prohibit modification and deletion)
func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	minH, maxH := math.MaxInt, 0
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x <= y {
			return y
		}
		return x
	}

	for _, w := range warehouse {
		minH, maxH = min(minH, w), max(maxH, w)
	}

	smallers, middles := make([]int, 0), make([]int, 0)
	for _, v := range boxes {
		if v <= minH {
			smallers = append(smallers, v)
		} else if v <= maxH {
			middles = append(middles, v)
		}
	}

	m, n := len(smallers), len(warehouse)
	if m >= n {
		return n
	}
	if len(middles) == 0 {
		return m
	}

	sort.Ints(middles)
	curIdx, res := n-1-m, m
	for curIdx > 0 {
		t, i := middles[0], 0
		middles = middles[1:]

		for i < curIdx && t <= warehouse[i] {
			i++
		}
		if i > 0 {
			warehouse[i-1] = t
			res++
		}
		curIdx = i - 1
		if len(middles) == 0 {
			break
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	boxes, warehouse := []int{4, 4, 1, 1}, []int{5, 4, 3, 3, 1}
	println(maxBoxesInWarehouse(boxes, warehouse))
}
