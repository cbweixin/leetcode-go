package main

//Given a string containing digits from 2-9 inclusive, return all possible lette
//r combinations that the number could represent. Return the answer in any order.
//
//
// A mapping of digit to letters (just like on the telephone buttons) is given b
//elow. Note that 1 does not map to any letters.
//
//
//
//
// Example 1:
//
//
//Input: digits = "23"
//Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// Example 2:
//
//
//Input: digits = ""
//Output: []
//
//
// Example 3:
//
//
//Input: digits = "2"
//Output: ["a","b","c"]
//
//
//
// Constraints:
//
//
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
//
// Related Topics String Backtracking
// 👍 4648 👎 452

// 2020-10-27 11:43:25

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 数字字母映射
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var ans []string
	var dfs func(int, string)

	// DFS函数定义
	dfs = func(i int, path string) {
		if i >= len(digits) {
			ans = append(ans, path)
			return
		}

		for _, c := range mp[string(digits[i])] {
			dfs(i+1, path+string(c))
		}
	}
	// 执行回溯函数
	dfs(0, "")
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
