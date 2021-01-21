package main

//In an array A containing only 0s and 1s, a K-bit flip consists of choosing a (
//contiguous) subarray of length K and simultaneously changing every 0 in the suba
//rray to 1, and every 1 in the subarray to 0.
//
// Return the minimum number of K-bit flips required so that there is no 0 in th
//e array. If it is not possible, return -1.
//
//
//
// Example 1:
//
//
//Input: A = [0,1,0], K = 1
//Output: 2
//Explanation: Flip A[0], then flip A[2].
//
//
//
// Example 2:
//
//
//Input: A = [1,1,0], K = 2
//Output: -1
//Explanation: No matter how we flip subarrays of size 2, we can't make the arra
//y become [1,1,1].
//
//
//
// Example 3:
//
//
//Input: A = [0,0,0,1,0,1,1,0], K = 3
//Output: 3
//Explanation:
//Flip A[0],A[1],A[2]: A becomes [1,1,1,1,0,1,1,0]
//Flip A[4],A[5],A[6]: A becomes [1,1,1,1,1,0,0,0]
//Flip A[5],A[6],A[7]: A becomes [1,1,1,1,1,1,1,1]
//
//
//
//
//
//
// Note:
//
//
// 1 <= A.length <= 30000
// 1 <= K <= A.length
// Related Topics Greedy Sliding Window
// 👍 454 👎 34

// 2021-01-21 09:29:40

//leetcode submit region begin(Prohibit modification and deletion)
func minKBitFlips(A []int, K int) int {
	res, L := 0, len(A)
	que := []int{}

	for i, v := range A {
		if len(que) > 0 && que[0]+K == i {
			que = que[1:]
		}
		if len(que)%2 == v {
			if L-i < K {
				return -1
			}
			que = append(que, i)
			res++
		}

	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
