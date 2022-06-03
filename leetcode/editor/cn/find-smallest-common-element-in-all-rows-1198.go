package main

//给你一个 m x n 的矩阵 mat，其中每一行的元素均符合 严格递增 。请返回 所有行中的 最小公共元素 。
//
// 如果矩阵中没有这样的公共元素，就请返回 -1。
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
//输出：5
//
//
// 示例 2:
//
//
//输入：mat = [[1,2,3],[2,3,4],[2,3,5]]
//输出： 2
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 500
// 1 <= mat[i][j] <= 10⁴
// mat[i] 已按严格递增顺序排列。
//
// Related Topics 数组 哈希表 二分查找 计数 矩阵 👍 25 👎 0

// 2022-06-02 21:42:58
//leetcode submit region begin(Prohibit modification and deletion)
func smallestCommonElement(mat [][]int) int {
	count := make(map[int]int)
	l, m := len(mat), len(mat[0])
	for i := 0; i < l; i++ {
		for j := 0; j < m; j++ {
			if _, ok := count[mat[i][j]]; !ok {
				count[mat[i][j]] = 1
			} else {
				count[mat[i][j]]++
			}
			if count[mat[i][j]] == l {
				return mat[i][j]
			}

		}
	}
	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
