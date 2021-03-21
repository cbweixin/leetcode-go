package main

//You have n super washing machines on a line. Initially, each washing machine h
//as some dresses or is empty.
//
//
// For each move, you could choose any m (1 ≤ m ≤ n) washing machines, and pass
//one dress of each washing machine to one of its adjacent washing machines at the
// same time .
//
// Given an integer array representing the number of dresses in each washing mac
//hine from left to right on the line, you should find the minimum number of moves
// to make all the washing machines have the same number of dresses. If it is not
//possible to do it, return -1.
//
// Example1
//
//Input: [1,0,5]
//
//Output: 3
//
//Explanation:
//1st move:    1     0 <-- 5    =>    1     1     4
//2nd move:    1 <-- 1 <-- 4    =>    2     1     3
//3rd move:    2     1 <-- 3    =>    2     2     2
//
//
// Example2
//
//Input: [0,3,0]
//
//Output: 2
//
//Explanation:
//1st move:    0 <-- 3     0    =>    1     2     0
//2nd move:    1     2 --> 0    =>    1     1     1
//
//
// Example3
//
//Input: [0,2,0]
//
//Output: -1
//
//Explanation:
//It's impossible to make all the three washing machines have the same number of
// dresses.
//
//
//
//
// Note:
//
// The range of n is [1, 10000].
// The range of dresses number in a super washing machine is [0, 1e5].
//
// Related Topics Math Dynamic Programming
// 👍 399 👎 161

// 2021-03-21 16:33:49

//leetcode submit region begin(Prohibit modification and deletion)
func findMinMoves(machines []int) int {
	s, l := 0, len(machines)
	for _, v := range machines {
		s += v
	}
	if s%l != 0 {
		return -1
	}
	avg := int(s / l)

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	max2 := func(x, y, z int) int {
		return max(max(x, y), z)
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	cnt, res := 0, 0
	for _, v := range machines {
		cnt += v - avg
		res = max2(res, abs(cnt), v-avg)
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
