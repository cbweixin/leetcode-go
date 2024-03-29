package main

import "strconv"

//给你一个非负整数 num ，返回它的「加密字符串」。
//
// 加密的过程是把一个整数用某个未知函数进行转化，你需要从下表推测出该转化函数：
//
//
//
//
//
// 示例 1：
//
// 输入：num = 23
//输出："1000"
//
//
// 示例 2：
//
// 输入：num = 107
//输出："101100"
//
//
//
//
// 提示：
//
//
// 0 <= num <= 10^9
//
// Related Topics 位运算 数学 字符串 👍 19 👎 0

// 2022-06-18 11:12:17
//leetcode submit region begin(Prohibit modification and deletion)
func encode(num int) string {
	return strconv.FormatInt(int64(num+1), 2)[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
