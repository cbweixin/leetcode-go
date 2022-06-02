package main

//一个坐标可以从 -infinity 延伸到 +infinity 的 无限大的 棋盘上，你的 骑士 驻扎在坐标为 [0, 0] 的方格里。
//
// 骑士的走法和中国象棋中的马相似，走 “日” 字：即先向左（或右）走 1 格，再向上（或下）走 2 格；或先向左（或右）走 2 格，再向上（或下）走 1 格
//。
//
// 每次移动，他都可以按图示八个方向之一前进。
//
//
//
// 返回 骑士前去征服坐标为 [x, y] 的部落所需的最小移动次数 。本题确保答案是一定存在的。
//
//
//
// 示例 1：
//
//
//输入：x = 2, y = 1
//输出：1
//解释：[0, 0] → [2, 1]
//
//
// 示例 2：
//
//
//输入：x = 5, y = 5
//输出：4
//解释：[0, 0] → [2, 1] → [4, 2] → [3, 4] → [5, 5]
//
//
//
//
// 提示：
//
//
// -300 <= x, y <= 300
// 0 <= |x| + |y| <= 300
//
// Related Topics 广度优先搜索 👍 73 👎 0

// 2022-05-31 06:40:25
//leetcode submit region begin(Prohibit modification and deletion)
func minKnightMoves(x int, y int) int {
	res := 0
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
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	x, y = min(abs(x), abs(y)), max(abs(x), abs(y))

	for x > 4 || y > 4 {
		res++
		if x >= y {
			x -= 2
			if y >= 1 {
				y -= 1
			} else {
				y += 1
			}
		} else {
			y -= 2
			if x >= 1 {
				x -= 1
			} else {
				x += 1
			}
		}
	}

	dirs := [][]int{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}
	que := [][]int{{0, 0, 0}}

	for len(que) > 0 {
		a := que[0]
		que = que[1:]
		row, col, steps := a[0], a[1], a[2]
		if row == x && col == y {
			return steps + res
		}
		for _, v := range dirs {
			dr, dc := v[0], v[1]
			if (x-row)*dr < 0 && (y-col)*dc < 0 {
				continue
			}
			que = append(que, []int{row + dr, col + dc, steps + 1})
		}
	}
	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	println(minKnightMoves(25, -87))
	println(minKnightMoves(270, -21))
}
