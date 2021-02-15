package main

//Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elemen
//ts appear twice and others appear once. 
//
// Find all the elements of [1, n] inclusive that do not appear in this array. 
//
// Could you do it without extra space and in O(n) runtime? You may assume the r
//eturned list does not count as extra space. 
//
// Example:
// 
//Input:
//[4,3,2,7,8,2,3,1]
//
//Output:
//[5,6]
// 
// Related Topics Array 
// 👍 3853 👎 281

// 2021-02-14 19:37:52

//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, v := range nums {
		idx := abs(v) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}

	}

	for k, v := range nums {
		if v > 0 {
			res = append(res, k+1)
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
