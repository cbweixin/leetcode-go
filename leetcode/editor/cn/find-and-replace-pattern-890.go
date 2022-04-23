package main

import "fmt"

// 你有一个单词列表 words 和一个模式 pattern，你想知道 words 中的哪些单词与模式匹配。
//
// 如果存在字母的排列 p ，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。
//
// （回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）
//
// 返回 words 中与给定模式匹配的单词列表。
//
// 你可以按任何顺序返回答案。
//
//
//
// 示例：
//
// 输入：words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
// 输出：["mee","aqq"]
// 解释：
// "mee" 与模式匹配，因为存在排列 {a -> m, b -> e, ...}。
// "ccc" 与模式不匹配，因为 {a -> c, b -> c, ...} 不是排列。
// 因为 a 和 b 映射到同一个字母。
//
//
//
// 提示：
//
//
// 1 <= words.length <= 50
// 1 <= pattern.length = words[i].length <= 20
//
// Related Topics 数组 哈希表 字符串 👍 125 👎 0

// 2022-04-23 15:49:35
// leetcode submit region begin(Prohibit modification and deletion)
func findAndReplacePattern(words []string, pattern string) []string {
	get_int := func(c int) int {
		return c - 'a'
	}
	match := func(w, p string) bool {
		lookup := make([]int, 26)
		lookup2 := make([]int, 26)
		for i := 0; i < 26; i++ {
			lookup[i] = -1
		}
		for i := 0; i < len(p); i++ {
			wIdx, pIdx := get_int(int(w[i])), get_int(int(p[i]))
			if lookup[wIdx] == -1 {
				lookup[wIdx] = pIdx
				lookup2[pIdx]++
			} else if lookup[wIdx] != pIdx {
				return false
			}
		}
		for i := 0; i < 26; i++ {
			if lookup2[i] > 1 {
				return false
			}
		}
		return true
	}
	var res []string
	for _, w := range words {
		if match(w, pattern) {
			res = append(res, w)
		}
	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"

	fmt.Println(findAndReplacePattern(words, pattern))
}
