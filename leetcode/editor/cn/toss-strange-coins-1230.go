package main

//有一些不规则的硬币。在这些硬币中，prob[i] 表示第 i 枚硬币正面朝上的概率。
//
// 请对每一枚硬币抛掷 一次，然后返回正面朝上的硬币数等于 target 的概率。
//
//
//
// 示例 1：
//
// 输入：prob = [0.4], target = 1
//输出：0.40000
//
//
// 示例 2：
//
// 输入：prob = [0.5,0.5,0.5,0.5,0.5], target = 0
//输出：0.03125
//
//
//
//
// 提示：
//
//
// 1 <= prob.length <= 1000
// 0 <= prob[i] <= 1
// 0 <= target <= prob.length
// 如果答案与标准答案的误差在 10^-5 内，则被视为正确答案。
//
// Related Topics 数学 动态规划 概率与统计 👍 56 👎 0

// 2022-06-12 19:46:49
//leetcode submit region begin(Prohibit modification and deletion)
// dp
func probabilityOfHeads(prob []float64, target int) float64 {
	l := len(prob)
	dp := make([]float64, l+1)
	dp[0] = 1.0
	for i := 1; i < l+1; i++ {
		for j := i; j >= 0; j-- {
			if j > 0 {
				dp[j] = dp[j]*(1-prob[i-1]) + dp[j-1]*prob[i-1]
			} else {
				dp[j] = dp[j] * (1 - prob[i-1])
			}
		}
	}
	return dp[target]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	prob := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	target := 0
	res := probabilityOfHeads(prob, target)
	println(res)
}
