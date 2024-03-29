package main

import "sort"

// 给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。
//
// 用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。
//
// 满足条件的二叉树一共有多少个？答案可能很大，返回 对 10⁹ + 7 取余 的结果。
//
//
//
// 示例 1:
//
//
// 输入: arr = [2, 4]
// 输出: 3
// 解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
//
// 示例 2:
//
//
// 输入: arr = [2, 4, 5, 10]
// 输出: 7
// 解释: 可以得到这些二叉树: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 1000
// 2 <= arr[i] <= 10⁹
// arr 中的所有值 互不相同
//
// Related Topics 数组 哈希表 动态规划 👍 78 👎 0

// 2022-03-19 21:00:39
// leetcode submit region begin(Prohibit modification and deletion)
// dp
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	l := len(arr)
	dp := make([]int, l)
	mod := 1000000007
	for i := 0; i < l; i++ {
		dp[i] = 1
	}
	for i, v := range arr {
		if i == 0 {
			continue
		}
		j, k := 0, i-1
		for j <= k {
			t := arr[j] * arr[k]
			if t < v {
				j++
			} else if t > arr[i] {
				k--
			} else {
				dp[i] = (dp[i] + dp[j]*dp[k]) % mod
				if arr[j] != arr[k] {
					dp[i] = (dp[i] + dp[j]*dp[k]) % mod
				}
				j++
				k--
			}
		}
		// dp[i] %= mod
	}

	res := 0
	for _, v := range dp {
		res = (res + v) % mod
	}
	return res

}

// leetcode submit region end(Prohibit modification and deletion)
