package main

import (
	"bytes"
	"sort"
	"strconv"
)

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
// 输入：nums = [10,2]
// 输出："210"
//
// 示例 2：
//
//
// 输入：nums = [3,30,34,5,9]
// 输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 贪心 字符串 排序 👍 1006 👎 0

// 2022-09-03 22:49:54
// leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	sort.Slice(
		nums, func(i, j int) bool {
			s1, s2 := strconv.Itoa(i), strconv.Itoa(j)
			return s1+s2 > s2+s1
		},
	)
	var b bytes.Buffer
	for _, v := range nums {
		b.WriteString(strconv.Itoa(v))
	}

	return b.String()
}

// leetcode submit region end(Prohibit modification and deletion)
