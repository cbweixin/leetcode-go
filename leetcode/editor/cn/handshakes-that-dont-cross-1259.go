package main

//偶数 个人站成一个圆，总人数为 num_people 。每个人与除自己外的一个人握手，所以总共会有 num_people / 2 次握手。
//
// 将握手的人之间连线，请你返回连线不会相交的握手方案数。
//
// 由于结果可能会很大，请你返回答案 模 10^9+7 后的结果。
//
//
//
// 示例 1：
//
// 输入：num_people = 2
//输出：1
//
//
// 示例 2：
//
//
//
// 输入：num_people = 4
//输出：2
//解释：总共有两种方案，第一种方案是 [(1,2),(3,4)] ，第二种方案是 [(2,3),(4,1)] 。
//
//
// 示例 3：
//
//
//
// 输入：num_people = 6
//输出：5
//
//
// 示例 4：
//
// 输入：num_people = 8
//输出：14
//
//
//
//
// 提示：
//
//
// 2 <= num_people <= 1000
// num_people % 2 == 0
//
// Related Topics 数学 动态规划 👍 25 👎 0

// 2022-06-20 13:22:27
//leetcode submit region begin(Prohibit modification and deletion)
func numberOfWays(numPeople int) int {
	n, mod := numPeople+1, 1000000007
	dp := make([]int, n, n)
	dp[0], dp[2] = 1, 1

	for i := 4; i < n; i += 2 {
		for j := 1; j < i; j += 2 {
			dp[i] = (dp[i] + (dp[j-1]*dp[i-j-1])%mod) % mod
		}
	}

	return dp[numPeople]
}

//leetcode submit region end(Prohibit modification and deletion)
