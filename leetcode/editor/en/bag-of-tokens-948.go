package main

import (
	"sort"
)

//You have an initial power of P, an initial score of 0, and a bag of tokens whe
//re tokens[i] is the value of the ith token (0-indexed).
//
// Your goal is to maximize your total score by potentially playing each token i
//n one of two ways:
//
//
// If your current power is at least tokens[i], you may play the ith token face
//up, losing tokens[i] power and gaining 1 score.
// If your current score is at least 1, you may play the ith token face down, ga
//ining tokens[i] power and losing 1 score.
//
//
// Each token may be played at most once and in any order. You do not have to pl
//ay all the tokens.
//
// Return the largest possible score you can achieve after playing any number of
// tokens.
//
//
// Example 1:
//
//
//Input: tokens = [100], P = 50
//Output: 0
//Explanation: Playing the only token in the bag is impossible because you eithe
//r have too little power or too little score.
//
//
// Example 2:
//
//
//Input: tokens = [100,200], P = 150
//Output: 1
//Explanation: Play the 0th token (100) face up, your power becomes 50 and score
// becomes 1.
//There is no need to play the 1st token since you cannot play it face up to add
// to your score.
//
//
// Example 3:
//
//
//Input: tokens = [100,200,300,400], P = 200
//Output: 2
//Explanation: Play the tokens in this order to get a score of 2:
//1. Play the 0th token (100) face up, your power becomes 100 and score becomes
//1.
//2. Play the 3rd token (400) face down, your power becomes 500 and score become
//s 0.
//3. Play the 1st token (200) face up, your power becomes 300 and score becomes
//1.
//4. Play the 2nd token (300) face up, your power becomes 0 and score becomes 2.
//
//
//
//
// Constraints:
//
//
// 0 <= tokens.length <= 1000
// 0 <= tokens[i], P < 104
//
// Related Topics Greedy
// 👍 400 👎 240

// 2020-10-25 11:27:59

//leetcode submit region begin(Prohibit modification and deletion)

func bagOfTokensScore(tokens []int, P int) int {
	if len(tokens) == 0 {
		return 0
	}
	sort.Ints(tokens)
	i, j := 0, len(tokens)-1
	ans, score := 0, 0
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for (i <= j && j >= 0) && (P >= 0 || score >= 0) {
		for i <= j && j >= 0 && P-tokens[i] >= 0 {
			P -= tokens[i]
			score++
			i++
		}
		ans = max(ans, score)
		if (i <= j && j >= 0) && score > 0 {
			P += tokens[j]
			score--
			j--
		}

	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	tokens := []int{33, 4, 28, 24, 96}
	P := 35

	println(bagOfTokensScore(tokens, P))
}
