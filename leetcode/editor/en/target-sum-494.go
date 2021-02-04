package main

//You are given a list of non-negative integers, a1, a2, ..., an, and a target,
//S. Now you have 2 symbols + and -. For each integer, you should choose one from
//+ and - as its new symbol.
//
// Find out how many ways to assign symbols to make sum of integers equal to tar
//get S.
//
// Example 1:
//
//
//Input: nums is [1, 1, 1, 1, 1], S is 3.
//Output: 5
//Explanation:
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//There are 5 ways to assign symbols to make the sum of nums be target 3.
//
//
//
// Constraints:
//
//
// The length of the given array is positive and will not exceed 20.
// The sum of elements in the given array will not exceed 1000.
// Your output answer is guaranteed to be fitted in a 32-bit integer.
//
// Related Topics Dynamic Programming Depth-first Search
// 👍 3643 👎 147

// 2021-02-04 10:36:27

//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, S int) int {
	lookup := make(map[int]int)
	lookup[0] = 1

	for i := 0; i < len(nums); i++ {
		temp := make(map[int]int)
		for k, v := range lookup {
			c, ok := temp[k+nums[i]]
			if ok {
				temp[k+nums[i]] = c + v
			} else {
				temp[k+nums[i]] = v

			}

			c2, ok2 := temp[k-nums[i]]
			if ok2 {
				temp[k-nums[i]] = c2 + v
			} else {
				temp[k-nums[i]] = v
			}
		}

		lookup = temp
	}

	t, ok := lookup[S]
	if ok {
		return t
	}

	return 0

}

//leetcode submit region end(Prohibit modification and deletion)
