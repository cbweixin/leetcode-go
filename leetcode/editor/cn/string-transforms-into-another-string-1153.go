package main

//给出两个长度相同的字符串 str1 和 str2。请你帮忙判断字符串 str1 能不能在 零次 或 多次 转化 后变成字符串 str2。
//
// 每一次转化时，你可以将 str1 中出现的 所有 相同字母变成其他 任何 小写英文字母。
//
// 只有在字符串 str1 能够通过上述方式顺利转化为字符串 str2 时才能返回 true 。
//
//
//
// 示例 1：
//
//
//输入：str1 = "aabcc", str2 = "ccdee"
//输出：true
//解释：将 'c' 变成 'e'，然后把 'b' 变成 'd'，接着再把 'a' 变成 'c'。注意，转化的顺序也很重要。
//
//
// 示例 2：
//
//
//输入：str1 = "leetcode", str2 = "codeleet"
//输出：false
//解释：我们没有办法能够把 str1 转化为 str2。
//
//
//
//
// 提示：
//
//
// 1 <= str1.length == str2.length <= 10⁴
// str1 和 str2 中都只会出现小写英文字母
//
// Related Topics 哈希表 字符串 👍 65 👎 0

// 2022-05-25 07:59:54
//leetcode submit region begin(Prohibit modification and deletion)
func canConvert(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s1_to_s2 := make(map[rune]rune)
	uniq := make(map[rune]bool)
	for _, c := range str2 {
		if _, ok := uniq[c]; !ok {
			uniq[c] = true
		}
	}

	if str1 != str2 && len(uniq) == 26 {
		return false
	}

	for i := 0; i < len(str1); i++ {
		_, ok := s1_to_s2[rune(str1[i])]
		if !ok {
			s1_to_s2[rune(str1[i])] = rune(str2[i])
		} else {
			if s1_to_s2[rune(str1[i])] != rune(str2[i]) {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
