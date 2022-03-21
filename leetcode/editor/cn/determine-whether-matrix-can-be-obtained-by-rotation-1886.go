package main

//给你两个大小为 n x n 的二进制矩阵 mat 和 target 。现 以 90 度顺时针轮转 矩阵 mat 中的元素 若干次 ，如果能够使 mat 与
//target 一致，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
//输出：true
//解释：顺时针轮转 90 度一次可以使 mat 和 target 一致。
//
//
// 示例 2：
//
//
//输入：mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
//输出：false
//解释：无法通过轮转矩阵中的元素使 equal 与 target 一致。
//
//
// 示例 3：
//
//
//输入：mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
//输出：true
//解释：顺时针轮转 90 度两次可以使 mat 和 target 一致。
//
//
//
//
// 提示：
//
//
// n == mat.length == target.length
// n == mat[i].length == target[i].length
// 1 <= n <= 10
// mat[i][j] 和 target[i][j] 不是 0 就是 1
//
// Related Topics 数组 矩阵 👍 17 👎 0

// 2022-03-21 09:21:29
//leetcode submit region begin(Prohibit modification and deletion)
func findRotation(mat [][]int, target [][]int) bool {
	r, c := len(mat), len(mat[0])
	// rotate 90 : (x,y) -> (y, n-x)
	// roate 180 : (x,y) -> (n-x, n-y)
	// rotate 270 : (x,y)->(n-y,x)

	b1, b2, b3, b4 := true, true, true, true
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] != target[i][j] {
				b1 = false
			}
			if mat[i][j] != target[j][c-1-i] {
				b2 = false
			}
			if mat[i][j] != target[c-1-i][c-1-j] {
				b3 = false
			}
			if mat[i][j] != target[c-1-j][i] {
				b4 = false
			}
		}
	}

	return b1 || b2 || b3 || b4
}

//leetcode submit region end(Prohibit modification and deletion)
