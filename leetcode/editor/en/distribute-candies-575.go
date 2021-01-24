package main

//Alice has n candies, where the ith candy is of type candyType[i]. Alice notice
//d that she started to gain weight, so she visited a doctor.
//
// The doctor advised Alice to only eat n / 2 of the candies she has (n is alway
//s even). Alice likes her candies very much, and she wants to eat the maximum num
//ber of different types of candies while still following the doctor's advice.
//
// Given the integer array candyType of length n, return the maximum number of d
//ifferent types of candies she can eat if she only eats n / 2 of them.
//
//
// Example 1:
//
//
//Input: candyType = [1,1,2,2,3,3]
//Output: 3
//Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 type
//s, she can eat one of each type.
//
//
// Example 2:
//
//
//Input: candyType = [1,1,2,3]
//Output: 2
//Explanation: Alice can only eat 4 / 2 = 2 candies. Whether she eats types [1,2
//], [1,3], or [2,3], she still can only eat 2 different types.
//
//
// Example 3:
//
//
//Input: candyType = [6,6,6,6]
//Output: 1
//Explanation: Alice can only eat 4 / 2 = 2 candies. Even though she can eat 2 c
//andies, she only has 1 type.
//
//
//
// Constraints:
//
//
// n == candyType.length
// 2 <= n <= 104
// n is even.
// -105 <= candyType[i] <= 105
//
// Related Topics Hash Table
// 👍 558 👎 922

// 2021-01-23 16:16:11

//leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, v := range candies {
		if _, ok := m[v]; !ok {
			m[v] = v
		}
	}
	if len(m) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(m)

}

//leetcode submit region end(Prohibit modification and deletion)
