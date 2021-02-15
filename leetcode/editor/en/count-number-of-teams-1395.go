package main

//There are n soldiers standing in a line. Each soldier is assigned a unique rat
//ing value.
//
// You have to form a team of 3 soldiers amongst them under the following rules:
//
//
//
// Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rat
//ing[k]).
// A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > ratin
//g[j] > rating[k]) where (0 <= i < j < k < n).
//
//
// Return the number of teams you can form given the conditions. (soldiers can b
//e part of multiple teams).
//
//
// Example 1:
//
//
//Input: rating = [2,5,3,4,1]
//Output: 3
//Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (
//5,3,1).
//
//
// Example 2:
//
//
//Input: rating = [2,1,3]
//Output: 0
//Explanation: We can't form any team given the conditions.
//
//
// Example 3:
//
//
//Input: rating = [1,2,3,4]
//Output: 4
//
//
//
// Constraints:
//
//
// n == rating.length
// 3 <= n <= 1000
// 1 <= rating[i] <= 105
// All the integers in rating are unique.
//
// Related Topics Array
// 👍 805 👎 114

// 2021-02-15 15:35:12

//leetcode submit region begin(Prohibit modification and deletion)
func numTeams(rating []int) int {
	l, res := len(rating), 0
	var b2i = map[bool]int8{false: 0, true: 1}

	if l < 3 {
		return 0
	}

	for i := 0; i < l; i++ {
		less, great := [2]int{0, 0}, [2]int{0, 0}
		for j := 0; j < l; j++ {
			if rating[i] > rating[j] {
				less[b2i[j < i]]++
			} else if rating[i] < rating[j] {
				great[b2i[j < i]]++
			}
		}
		res += less[0]*great[1] + less[1]*great[0]
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
