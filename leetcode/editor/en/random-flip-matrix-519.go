package main

import "math/rand"

//You are given the number of rows n_rows and number of columns n_cols of a 2D b
//inary matrix where all values are initially 0. Write a function flip which choos
//es a 0 value uniformly at random, changes it to 1, and then returns the position
// [row.id, col.id] of that value. Also, write a function reset which sets all val
//ues back to 0. Try to minimize the number of calls to system's Math.random() and
// optimize the time and space complexity.
//
// Note:
//
//
// 1 <= n_rows, n_cols <= 10000
// 0 <= row.id < n_rows and 0 <= col.id < n_cols
// flip will not be called when the matrix has no 0 values left.
// the total number of calls to flip and reset will not exceed 1000.
//
//
// Example 1:
//
//
//Input:
//["Solution","flip","flip","flip","flip"]
//[[2,3],[],[],[],[]]
//Output: [null,[0,1],[1,2],[1,0],[1,1]]
//
//
//
// Example 2:
//
//
//Input:
//["Solution","flip","flip","reset","flip"]
//[[1,2],[],[],[],[]]
//Output: [null,[0,0],[0,1],null,[0,0]]
//
//
// Explanation of Input Syntax:
//
// The input is two lists: the subroutines called and their arguments. Solution'
//s constructor has two arguments, n_rows and n_cols. flip and reset have no argum
//ents. Arguments are always wrapped with a list, even if there aren't any.
// Related Topics Random
// 👍 212 👎 69

// 2021-01-21 10:18:17

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	flipped map[int]int
	rows    int
	cols    int
	size    int
}

func Constructor(n_rows int, n_cols int) Solution {
	flip := make(map[int]int)
	size := n_rows * n_cols

	return Solution{flip, n_rows, n_cols, size}

}

func (this *Solution) Flip() []int {
	val := rand.Intn(this.size)
	this.size--
	x := val
	if v, ok := this.flipped[val]; ok {
		x = v
	}
	this.flipped[val] = this.size
	if v, ok := this.flipped[this.size]; ok {
		this.flipped[val] = v
	}
	return []int{int(x / this.cols), x % this.cols}

}

func (this *Solution) Reset() {
	this.flipped = make(map[int]int)
	this.size = this.rows * this.cols

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
//leetcode submit region end(Prohibit modification and deletion)
