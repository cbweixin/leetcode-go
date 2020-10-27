package main

import "strings"

//Given a pattern and a string s, find if s follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between a lett
//er in pattern and a non-empty word in s.
//
//
// Example 1:
//
//
//Input: pattern = "abba", s = "dog cat cat dog"
//Output: true
//
//
// Example 2:
//
//
//Input: pattern = "abba", s = "dog cat cat fish"
//Output: false
//
//
// Example 3:
//
//
//Input: pattern = "aaaa", s = "dog cat cat dog"
//Output: false
//
//
// Example 4:
//
//
//Input: pattern = "abba", s = "dog dog dog dog"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= pattern.length <= 300
// pattern contains only lower-case English letters.
// 1 <= s.length <= 3000
// s contains only lower-case English letters and spaces ' '.
// s does not contain any leading or trailing spaces.
// All the words in s are separated by a single space.
//
// Related Topics Hash Table
// 👍 1515 👎 192

// 2020-10-27 09:36:16

//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	p2s := make(map[byte]string)
	s2p := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		if w, ok := p2s[pattern[i]]; ok && w != words[i] {
			return false
		}
		if p, ok := s2p[words[i]]; ok && p != pattern[i] {
			return false
		}
		p2s[pattern[i]] = words[i]
		s2p[words[i]] = pattern[i]
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)
