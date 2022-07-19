package main

// 给定一个字符串列表 dict ，其中所有字符串的长度都相同。
//
// 当存在两个字符串在相同索引处只有一个字符不同时，返回 True ，否则返回 False 。
//
//
//
// 示例 1：
//
//
// 输入：dict = ["abcd","acbd", "aacd"]
// 输出：true
// 解释：字符串 "abcd" 和 "aacd" 只在索引 1 处有一个不同的字符。
//
//
// 示例 2：
//
//
// 输入：dict = ["ab","cd","yz"]
// 输出：false
//
//
// 示例 3：
//
//
// 输入：dict = ["abcd","cccc","abyd","abab"]
// 输出：true
//
//
//
//
// 提示：
//
//
// dict 中的字符数小于或等于 10^5 。
// dict[i].length == dict[j].length
// dict[i] 是互不相同的。
// dict[i] 只包含小写英文字母。
//
//
//
//
// 进阶：你可以以 O(n*m) 的复杂度解决问题吗？其中 n 是列表 dict 的长度，m 是字符串的长度。
// Related Topics 哈希表 字符串 哈希函数 滚动哈希 👍 20 👎 0

// 2022-07-19 10:11:35
// leetcode submit region begin(Prohibit modification and deletion)
func differByOne(dict []string) bool {
	mod := 100000000007
	l, m := len(dict), len(dict[0])
	hash := make([]int, l)

	for i := 0; i < l; i++ {
		for j := 0; j < m; j++ {
			// rune is a alias of int32
			hash[i] = (26*hash[i] + int(rune(dict[i][j])-rune('a'))) % mod
		}
	}

	seen, base := make(map[int]struct{}), 1

	for j := m - 1; j >= 0; j-- {
		for i := 0; i < l; i++ {
			newHash := (hash[i] - base*int(rune(dict[i][j])-rune('a'))) % mod
			if _, ok := seen[newHash]; ok {
				return true
			}
			seen[newHash] = struct{}{}
		}
		base = (base * 26) % mod
	}

	return false

}

// leetcode submit region end(Prohibit modification and deletion)
