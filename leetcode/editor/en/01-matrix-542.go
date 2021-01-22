package main

//Given a matrix consists of 0 and 1, find the distance of the nearest 0 for eac
//h cell.
//
// The distance between two adjacent cells is 1.
//
//
//
// Example 1:
//
//
//Input:
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
//
//Output:
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
//
//
// Example 2:
//
//
//Input:
//[[0,0,0],
// [0,1,0],
// [1,1,1]]
//
//Output:
//[[0,0,0],
// [0,1,0],
// [1,2,1]]
//
//
//
//
// Note:
//
//
// The number of elements of the given matrix will not exceed 10,000.
// There are at least one 0 in the given matrix.
// The cells are adjacent in only four directions: up, down, left and right.
//
// Related Topics Depth-first Search Breadth-first Search
// 👍 2036 👎 121

// 2021-01-22 10:34:19

//leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(matrix [][]int) [][]int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	x, y := len(matrix), len(matrix[0])
	que := [][]int{}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if matrix[i][j] == 0 {
				que = append(que, []int{i, j})
			} else {
				matrix[i][j] = 100001
			}
		}
	}

	for len(que) > 0 {
		tmp := que
		que = nil
		for _, val := range tmp {
			i, j := val[0], val[1]
			for _, dir := range dirs {
				n_i, n_j := i+dir[0], j+dir[1]
				if n_i >= 0 && n_i < x && n_j >= 0 && n_j < y && matrix[n_i][n_j] > matrix[i][j] {
					matrix[n_i][n_j] = matrix[i][j] + 1
					que = append(que, []int{n_i, n_j})
				}
			}
		}
	}

	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
