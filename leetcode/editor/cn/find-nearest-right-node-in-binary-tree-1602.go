package main

// 给定一棵二叉树的根节点 root 和树中的一个节点 u ，返回与 u 所在层中距离最近的右侧节点，当 u 是所在层中最右侧的节点，返回 null 。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [1,2,3,null,4,5,6], u = 4
// 输出：5
// 解释：节点 4 所在层中，最近的右侧节点是节点 5。
//
//
// 示例 2：
//
//
//
//
// 输入：root = [3,null,4,2], u = 2
// 输出：null
// 解释：2 的右侧没有节点。
//
//
// 示例 3：
//
//
// 输入：root = [1], u = 1
// 输出：null
//
//
// 示例 4：
//
//
// 输入：root = [3,4,2,null,null,null,1], u = 4
// 输出：2
//
//
//
//
// 提示:
//
//
// 树中节点个数的范围是 [1, 10⁵] 。
// 1 <= Node.val <= 10⁵
// 树中所有节点的值是唯一的。
// u 是以 root 为根的二叉树的一个节点。
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 9 👎 0

// 2022-08-02 15:04:19
// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type LevelPair struct {
	Level int
	Node  *TreeNode
}

func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	que := []LevelPair{{
		Level: 0,
		Node:  root,
	}}

	for len(que) > 0 {
		lp := que[0]
		que = que[1:]
		level, node := lp.Level, lp.Node
		if node == u {
			if len(que) > 0 && que[0].Level == level {
				return que[0].Node
			}
			return nil
		}
		if node.Left != nil {
			que = append(que, LevelPair{
				Level: level + 1,
				Node:  node.Left,
			})
		}
		if node.Right != nil {
			que = append(que, LevelPair{
				Level: level + 1,
				Node:  node.Right,
			})
		}
	}

	return nil

}

// leetcode submit region end(Prohibit modification and deletion)
