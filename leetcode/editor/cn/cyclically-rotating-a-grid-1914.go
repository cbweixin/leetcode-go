package main

import (
	"fmt"
	"reflect"
)

// 给你一个大小为 m x n 的整数矩阵 grid ，其中 m 和 n 都是 偶数 ；另给你一个整数 k 。
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
// 输出：[[10,20],[40,30]]
// 解释：上图展示了矩阵在执行循环轮转操作时每一步的状态。
//
// 示例 2：
//
//
// 输入：grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
// 输出：[[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
// 解释：上图展示了矩阵在执行循环轮转操作时每一步的状态。
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
// leetcode submit region begin(Prohibit modification and deletion)
func rotateGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	layerDepth := min(r, c) >> 1
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
		m := k % len(arr)
		// arr[:k]... means unpack slice arr[:k]
		arr = append(arr[m:], arr[:m]...)
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

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	// var grid = [][]int{{40, 10}, {30, 20}}
	// fmt.Println(rotateGrid(grid, 1))
	// grid = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// fmt.Println(rotateGrid(grid, 2))

	var grid = [][]int{
		{3970, 1906, 3608, 298, 3072, 3546, 1502, 773, 4388, 3115, 747, 3937},
		{2822, 304, 4179, 1780, 1709, 1058, 3645, 681, 2910, 2513, 4357, 1038},
		{4471, 2443, 218, 550, 2766, 4780, 1997, 1672, 4095, 161, 4645, 3838},
		{2035, 2350, 3653, 4127, 3208, 4717, 4347, 3452, 1601, 3725, 3060, 2270},
		{188, 2278, 81, 3454, 3204, 1897, 2862, 4381, 3704, 2587, 743, 3832},
		{996, 4499, 66, 2742, 1761, 1189, 608, 509, 2344, 3271, 3076, 108},
		{3274, 2042, 2157, 3226, 2938, 3766, 2610, 4510, 219, 1276, 3712, 4143},
		{744, 234, 2159, 4478, 4161, 4549, 4214, 4272, 701, 4376, 3110, 4896},
		{4431, 1011, 757, 2690, 83, 3546, 946, 1122, 2216, 3944, 2715, 2842},
		{898, 4087, 703, 4153, 3297, 2968, 3268, 4717, 1922, 2527, 3139, 1516},
		{1086, 1090, 302, 1273, 2292, 234, 3268, 2284, 4203, 3838, 2227, 3651},
		{2055, 4406, 2278, 3351, 3217, 2506, 4525, 233, 3829, 63, 4470, 3170},
		{3797, 3276, 1755, 1727, 1131, 4108, 3633, 1835, 1345, 1293, 2778, 2805},
		{1215, 84, 282, 2721, 2360, 2321, 1435, 2617, 1202, 2876, 3420, 3034},
	}

	arr := rotateGrid(grid, 405548684)
	fmt.Println(arr)
	expected := [][]int{
		{188, 2035, 4471, 2822, 3970, 1906, 3608, 298, 3072, 3546, 1502, 773},
		{996, 1058, 3645, 681, 2910, 2513, 4357, 4645, 3060, 743, 3076, 4388},
		{3274, 1709, 4376, 3944, 2527, 3838, 63, 3829, 233, 4525, 3712, 3115},
		{744, 1780, 1276, 4478, 3226, 2742, 3454, 4127, 3208, 2506, 3110, 747},
		{4431, 4179, 3271, 2690, 83, 4161, 2938, 1761, 4717, 3217, 2715, 3937},
		{898, 304, 2587, 4153, 3297, 946, 3546, 3204, 4347, 3351, 3139, 1038},
		{1086, 2443, 3725, 1273, 2968, 4214, 4549, 1897, 3452, 2278, 2227, 3838},
		{2055, 2350, 161, 2292, 3268, 2610, 3766, 2862, 1601, 302, 4470, 2270},
		{3797, 2278, 4095, 234, 4717, 608, 1189, 4381, 3704, 703, 2778, 3832},
		{1215, 4499, 1672, 3268, 1122, 4272, 4510, 509, 2344, 757, 1293, 108},
		{84, 2042, 1997, 2284, 4203, 1922, 2216, 701, 219, 2159, 1345, 4143},
		{282, 234, 4780, 2766, 550, 218, 3653, 81, 66, 2157, 1835, 4896},
		{2721, 1011, 4087, 1090, 4406, 3276, 1755, 1727, 1131, 4108, 3633, 2842},
		{2360, 2321, 1435, 2617, 1202, 2876, 3420, 3034, 2805, 3170, 3651, 1516},
	}

	a := reflect.DeepEqual(arr, expected)

	fmt.Println(a)

}
