package main

//In a array A of size 2N, there are N+1 unique elements, and exactly one of the
//se elements is repeated N times.
//
// Return the element repeated N times.
//
//
//
//
//
//
//
// Example 1:
//
//
//Input: [1,2,3,3]
//Output: 3
//
//
//
// Example 2:
//
//
//Input: [2,1,2,5,3,2]
//Output: 2
//
//
//
// Example 3:
//
//
//Input: [5,1,5,2,5,3,5,4]
//Output: 5
//
//
//
//
// Note:
//
//
// 4 <= A.length <= 10000
// 0 <= A[i] < 10000
// A.length is even
//
//
//
//
// Related Topics Hash Table
// 👍 527 👎 246

// 2020-10-26 14:15:20

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedNTimes1(A []int) int {
	// knowledge, use map as set
	set := make(map[int]bool)

	for _, a := range A {
		// review , check existence of an element in map
		if _, ok := set[a]; ok {
			return a
		}
		set[a] = true
	}
	return -1

}

func repeatedNTimes(A []int) int {
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] || A[i] == A[i-2] {
			return A[i]
		}
	}

	return A[0]

}

//leetcode submit region end(Prohibit modification and deletion)
