package main

//斐波那契树是一种按这种规则函数 order(n) 创建的二叉树：
//
//
// order(0) 是空树。
// order(1) 是一棵只有一个节点的二叉树。
// order(n) 是一棵根节点的左子树为 order(n - 2) 、右子树为 order(n - 1) 的二叉树。
//
//
// Alice 和 Bob 在玩一种关于斐波那契树的游戏，由 Alice 先手。在每个回合中，每个玩家选择一个节点，然后移除该节点及其子树。只能删除根节点
//root 的玩家输掉这场游戏。
//
// 给定一个整数 n，假定两名玩家都按最优策略进行游戏，若 Alice 赢得这场游戏，返回 true 。若 Bob 赢得这场游戏，返回 false 。
//
// 一棵二叉树的子树 tree 是由 tree 中某个节点及其所有后代节点组成的树。树 tree 也可当作自身的子树。
//
//
//
// 示例 1:
//
//
//输入: n = 3
//输出: true
//解释:
//Alice 移除右子树中的节点 1。
//Bob 要么移除左子树中的 1，要么移除右子树中的 2。
//Alice 可以移除 Bob 未移除的任意节点。
//Bob 只能删除根节点 3，所以 Bob 输了。
//返回 true，因为 Alice 赢了。
//
//
// 示例 2:
//
//
//输入: n = 1
//输出: false
//解释:
//Alice 只能移除根节点 1, 所以 Alice 输了。
//返回 false，因为 Alice 输了。
//
//
// 示例 3:
//
//
//输入: n = 2
//输出: true
//解释:
//Alice 删除节点 1.
//Bob 只能删除根节点 2，所以 Bob 输了。
//返回 true，因为 Alice 赢了。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
//
//
// Related Topics 树 数学 动态规划 二叉树 博弈 👍 3 👎 0

//2023-01-26 15:26:25

// leetcode submit region begin(Prohibit modification and deletion)
func findGameWinner(n int) bool {
	return n%6 != 1

}

//leetcode submit region end(Prohibit modification and deletion)
