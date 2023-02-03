package main

import "fmt"

//把字符串 s 看作 "abcdefghijklmnopqrstuvwxyz" 的无限环绕字符串，所以 s 看起来是这样的：
//
//
// "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." 。
//
//
// 现在给定另一个字符串 p 。返回 s 中 不同 的 p 的 非空子串 的数量 。
//
//
//
// 示例 1：
//
//
//输入：p = "a"
//输出：1
//解释：字符串 s 中只有 p 的一个 "a" 子字符。
//
//
// 示例 2：
//
//
//输入：p = "cac"
//输出：2
//解释：字符串 s 中只有 p 的两个子串 ("a", "c") 。
//
//
// 示例 3：
//
//
//输入：p = "zab"
//输出：6
//解释：在字符串 s 中有 p 的六个子串 ("z", "a", "b", "za", "ab", "zab") 。
//
//
//
//
// 提示：
//
//
// 1 <= p.length <= 10⁵
// p 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 367 👎 0

//2023-02-03 14:49:24

// leetcode submit region begin(Prohibit modification and deletion)
func findSubstringInWraproundString(p string) int {
	chars := make(map[byte]int, 26)
	for _, c := range p {
		chars[byte(c)-byte('a')] = 1
	}
	isConsecutive := func(x, y byte) bool {
		return (x-byte('a')+1)%26 == y-byte('a')
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	cur := 1
	l := len(p)
	for i := 1; i < l; i++ {
		if isConsecutive(p[i-1], p[i]) {
			cur += 1
		} else {
			cur = 1
		}
		chars[p[i]-byte('a')] = max(chars[p[i]-byte('a')], cur)
	}
	res := 0
	for _, v := range chars {
		res += v
	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findSubstringInWraproundString("zab"))
	//fmt.Println(byte('a'))
	//fmt.Println(byte('z'))
}
