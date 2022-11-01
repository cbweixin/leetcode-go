package main

import "math"

// 有一个花园，有 n 朵花，这些花都有一个用整数表示的美观度。这些花被种在一条线上。给定一个长度为 n 的整数类型数组 flowers ，每一个
// flowers[i] 表示第 i 朵花的美观度。
//
// 一个花园满足下列条件时，该花园是有效的。
//
//
// 花园中至少包含两朵花。
// 第一朵花和最后一朵花的美观度相同。
//
//
// 作为一个被钦定的园丁，你可以从花园中去除任意朵花（也可以不去除任意一朵）。你想要通过一种方法移除某些花朵，使得剩下的花园变得有效。花园的美观度是其中所有剩
// 余的花朵美观度之和。
//
// 返回你去除了任意朵花（也可以不去除任意一朵）之后形成的有效花园中最大可能的美观度。
//
//
//
// 示例 1：
//
// 输入: flowers = [1,2,3,1,2]
// 输出: 8
// 解释: 你可以修整为有效花园 [2,3,1,2] 来达到总美观度 2 + 3 + 1 + 2 = 8。
//
// 示例 2：
//
// 输入: flowers = [100,1,1,-3,1]
// 输出: 3
// 解释: 你可以修整为有效花园 [1,1,1] 来达到总美观度 1 + 1 + 1 = 3。
//
//
// 示例 3：
//
// 输入: flowers = [-1,-2,0,-1]
// 输出: -2
// 解释: 你可以修整为有效花园 [-1,-1] 来达到总美观度 -1 + -1 = -2。
//
//
//
//
// 提示：
//
//
// 2 <= flowers.length <= 10⁵
// -10⁴ <= flowers[i] <= 10⁴
// 去除一些花朵（可能没有）后，是有可能形成一个有效花园的。
//
//
// Related Topics 贪心 数组 前缀和 👍 4 👎 0

// 2022-10-31 22:28:27
// leetcode submit region begin(Prohibit modification and deletion)
func maximumBeauty(flowers []int) int {
	l := len(flowers)
	preSums := make([]int, l+1)
	positions := make(map[int]int)
	res := math.MinInt32

	max := func(x, y int) int {
		if x <= y {
			return y
		}
		return x
	}

	for i, v := range flowers {
		preSums[i+1] = preSums[i]
		if p, ok := positions[v]; ok {
			res = max(res, preSums[i]-preSums[p+1]+(v<<1))
		} else {
			positions[v] = i
		}
		if v > 0 {
			preSums[i+1] += v
		}
	}

	return res

}

// leetcode submit region end(Prohibit modification and deletion)
