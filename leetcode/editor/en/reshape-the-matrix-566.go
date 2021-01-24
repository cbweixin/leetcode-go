package main

//In MATLAB, there is a very useful function called 'reshape', which can reshape
// a matrix into a new one with different size but keep its original data.
//
//
//
//You're given a matrix represented by a two-dimensional array, and two positive
// integers r and c representing the row number and column number of the wanted re
//shaped matrix, respectively.
//
// The reshaped matrix need to be filled with all the elements of the original m
//atrix in the same row-traversing order as they were.
//
//
//
//If the 'reshape' operation with given parameters is possible and legal, output
// the new reshaped matrix; Otherwise, output the original matrix.
//
//
// Example 1:
//
//Input:
//nums =
//[[1,2],
// [3,4]]
//r = 1, c = 4
//Output:
//[[1,2,3,4]]
//Explanation: The row-traversing of nums is [1,2,3,4]. The new reshaped matrix
//is a 1 * 4 matrix, fill it row by row by using the previous list.
//
//
//
// Example 2:
//
//Input:
//nums =
//[[1,2],
// [3,4]]
//r = 2, c = 4
//Output:
//[[1,2],
// [3,4]]
//Explanation: There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So o
//utput the original matrix.
//
//
//
// Note:
//
// The height and width of the given matrix is in range [1, 100].
// The given r and c are all positive.
//
// Related Topics Array
// 👍 926 👎 106

// 2021-01-23 15:45:15

//leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	// review, how to create a 2-dimension slice
	arr := make([][]int, r)
	for i := range arr {
		arr[i] = make([]int, c)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := i*n + j
			b, d := int(a/c), a%c
			arr[b][d] = nums[i][j]
		}
	}

	return arr

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	matrixReshape(nums, 1, 4)
}
