package main

//给你一个大小为 m x n 的整数矩阵 grid ，其中 m 和 n 都是 偶数 ；另给你一个整数 k 。
//
// 矩阵由若干层组成，如下图所示，每种颜色代表一层：
//
//
//
// 矩阵的循环轮转是通过分别循环轮转矩阵中的每一层完成的。在对某一层进行一次循环旋转操作时，层中的每一个元素将会取代其 逆时针 方向的相邻元素。轮转示例如下：
//
//
// 返回执行 k 次循环轮转操作后的矩阵。
//
//
//
// 示例 1：
//
// 输入：grid = [[40,10],[30,20]], k = 1
//输出：[[10,20],[40,30]]
//解释：上图展示了矩阵在执行循环轮转操作时每一步的状态。
//
// 示例 2：
//
//
// 输入：grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
//输出：[[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
//解释：上图展示了矩阵在执行循环轮转操作时每一步的状态。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 50
// m 和 n 都是 偶数
// 1 <= grid[i][j] <= 5000
// 1 <= k <= 10⁹
//
// Related Topics 数组 矩阵 模拟 👍 11 👎 0

// 2022-02-15 21:29:19
//leetcode submit region begin(Prohibit modification and deletion)
func rotateGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	layerDepth := min(r, c) << 1
	for layer := 0; layer < layerDepth; layer++ {
		var arr []int
		// clockwise record the element in arr
		// up horizontal, from left to right
		for j := layer; j < c-layer; j++ {
			arr = append(arr, grid[layer][j])
		}
		// right vertical, from top to bottom
		for i := layer + 1; i < r-layer; i++ {
			arr = append(arr, grid[i][c-layer-1])
		}
		// bottom horizontal, from right to left
		for j := c - layer - 2; j >= layer; j-- {
			arr = append(arr, grid[r-layer-1][j])
		}
		// left vertical, from bottom to top
		for i := r - layer - 2; i > layer; i-- {
			arr = append(arr, grid[i][layer])
		}
		// left rotation
		k = k % len(arr)
		arr = append(arr[k:], arr[:k]...)
		idx := 0
		// clockwise record the element in arr
		// up horizontal, from left to right
		for j := layer; j < c-layer; j++ {
			grid[layer][j] = arr[idx]
			idx++
		}
		// right vertical, from top to bottom
		for i := layer + 1; i < r-layer; i++ {
			grid[i][c-layer-1] = arr[idx]
			idx++
		}
		// bottom horizontal, from right to left
		for j := c - layer - 2; j >= layer; j-- {
			grid[r-layer-1][j] = arr[idx]
			idx++
		}
		// left vertical, from bottom to top
		for i := r - layer - 2; i > layer; i-- {
			grid[i][layer] = arr[idx]
			idx++
		}

	}
	return grid
}

//leetcode submit region end(Prohibit modification and deletion)
