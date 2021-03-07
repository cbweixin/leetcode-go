package main

import (
	"math"
)

//There are several cards arranged in a row, and each card has an associated num
//ber of points The points are given in the integer array cardPoints.
//
// In one step, you can take one card from the beginning or from the end of the
//row. You have to take exactly k cards.
//
// Your score is the sum of the points of the cards you have taken.
//
// Given the integer array cardPoints and the integer k, return the maximum scor
//e you can obtain.
//
//
// Example 1:
//
//
//Input: cardPoints = [1,2,3,4,5,6,1], k = 3
//Output: 12
//Explanation: After the first step, your score will always be 1. However, choos
//ing the rightmost card first will maximize your total score. The optimal strateg
//y is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 1
//2.
//
//
// Example 2:
//
//
//Input: cardPoints = [2,2,2], k = 2
//Output: 4
//Explanation: Regardless of which two cards you take, your score will always be
// 4.
//
//
// Example 3:
//
//
//Input: cardPoints = [9,7,7,9,7,7,9], k = 7
//Output: 55
//Explanation: You have to take all the cards. Your score is the sum of points o
//f all cards.
//
//
// Example 4:
//
//
//Input: cardPoints = [1,1000,1], k = 1
//Output: 1
//Explanation: You cannot take the card in the middle. Your best score is 1.
//
//
// Example 5:
//
//
//Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
//Output: 202
//
//
//
// Constraints:
//
//
// 1 <= cardPoints.length <= 10^5
// 1 <= cardPoints[i] <= 10^4
// 1 <= k <= cardPoints.length
//
// Related Topics Array Dynamic Programming Sliding Window
// 👍 1130 👎 54

// 2021-03-07 14:24:01

//leetcode submit region begin(Prohibit modification and deletion)
func maxScore(cardPoints []int, k int) int {
	// sliding window
	l := len(cardPoints) - k
	left, right := 0, l
	res := math.MaxInt32

	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	arr_s := func(ll, rr int) int {
		ret := 0
		for i := ll; i < rr; i++ {
			ret += cardPoints[i]
		}
		return ret
	}
	s := arr_s(left, right)

	for right < len(cardPoints) {
		res = min(res, s)
		s += cardPoints[right] - cardPoints[left]
		left++
	}

	return arr_s(0, len(cardPoints)) - res

}

//leetcode submit region end(Prohibit modification and deletion)
