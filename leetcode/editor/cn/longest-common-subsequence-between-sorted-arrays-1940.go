package main

//给定一个由整数数组组成的数组arrays，其中arrays[i]是严格递增排序的，返回一个表示所有数组之间的最长公共子序列的整数数组。
//
// 子序列是从另一个序列派生出来的序列，删除一些元素或不删除任何元素，而不改变其余元素的顺序。
//
// 示例1:
//
//
//输入: arrays = [[1,3,4],
//               [1,4,7,9]]
//输出: [1,4]
//解释: 这两个数组中的最长子序列是[1,4]。
//
//
// 示例 2:
//
//
//输入: arrays = [[2,3,6,8],
//               [1,2,3,5,6,7,10],
//               [2,3,4,6,9]]
//输出: [2,3,6]
//解释: 这三个数组中的最长子序列是[2,3,6]。
//
//
// 示例 3:
//
//
//输入: arrays = [[1,2,3,4,5],
//               [6,7,8]]
//输出: []
//解释: 这两个数组之间没有公共子序列。
//
//
//
//
// 限制条件:
//
//
// 2 <= arrays.length <= 100
// 1 <= arrays[i].length <= 100
// 1 <= arrays[i][j] <= 100
// arrays[i] 是严格递增排序.
//
//
// Related Topics 数组 哈希表 计数 👍 5 👎 0

// 2023-01-01 18:16:09
//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(arrays [][]int) []int {
	freq := make(map[int]int, 100*100)
	for _, arr := range arrays {
		for _, v := range arr {
			freq[v] += 1
		}
	}
	res := make([]int, 0)
	for _, v := range arrays[0] {
		if freq[v] == len(arrays) {
			res = append(res, v)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
