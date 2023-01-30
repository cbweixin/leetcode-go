package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 2613 👎 0

//2023-01-29 21:05:24

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return s[:i]
			}
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
