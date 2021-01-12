package main

//Given an array nums of distinct integers, return all the possible permutations
//. You can return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:
// Input: nums = [0,1]
//Output: [[0,1],[1,0]]
// Example 3:
// Input: nums = [1]
//Output: [[1]]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.
//
// Related Topics Backtracking
// 👍 5180 👎 122

// 2021-01-11 22:09:40

//leetcode submit region begin(Prohibit modification and deletion)
func permute2(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

// We use a pointer for the result so we don't need to worry returning it.
func permuteRec(currComb, left []int, res *[][]int) {
	// We know that we found a new combination when we have no elements left.
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	// For the next iteration we consider all the left elements but the current one (idx).
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...), // Make sure to allocate a new slice.
			res,
		)
	}
}

func permute(nums []int) [][]int {
	c, done := make(chan []int, len(nums)), make(chan struct{})
	go p(nums, c, done)
	go func(done <-chan struct{}) {
		<-done
		close(c)
		return
	}(done)

	var r [][]int
	for e := range c {
		r = append(r, e)
	}
	return r
}

func p(a []int, c chan<- []int, done chan<- struct{}) {
	perm(a, 0, c)
	done <- struct{}{}
}

func perm(a []int, i int, c chan<- []int) {
	if i > len(a) {
		b := make([]int, len(a))
		copy(b, a)
		c <- b
		return
	}
	perm(a, i+1, c)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1, c)
		a[i], a[j] = a[j], a[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
