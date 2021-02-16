package main

//
//There is a strange printer with the following two special requirements:
//
//
// The printer can only print a sequence of the same character each time.
// At each turn, the printer can print new characters starting from and ending a
//t any places, and will cover the original existing characters.
//
//
//
//
//
//Given a string consists of lower English letters only, your job is to count th
//e minimum number of turns the printer needed in order to print it.
//
//
// Example 1:
//
//Input: "aaabbb"
//Output: 2
//Explanation: Print "aaa" first and then print "bbb".
//
//
//
// Example 2:
//
//Input: "aba"
//Output: 2
//Explanation: Print "aaa" first and then print "b" from the second place of the
// string, which will cover the existing character 'a'.
//
//
//
// Hint: Length of the given string will not exceed 100. Related Topics Dynamic
//Programming Depth-first Search
// 👍 543 👎 54

// 2021-02-16 11:47:49

//leetcode submit region begin(Prohibit modification and deletion)

type Key2 struct {
	X, Y int
}

func strangePrinter(s string) int {
	l := len(s)
	memo := make(map[Key2]int)
	min := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if v, ok := memo[Key2{i, j}]; ok {
			return v
		}
		if i > j {
			return 0
		}
		res := dfs(i+1, j) + 1
		for k := i + 1; k <= j; k++ {
			if s[i] == s[k] {
				res = min(res, dfs(i, k-1)+dfs(k+1, j))
			}
		}

		memo[Key2{i, j}] = res
		return res

	}

	return dfs(0, l-1)

}

//leetcode submit region end(Prohibit modification and deletion)
