package main

//给你一个二叉树，树中每个节点都含有一个附加的随机指针，该指针可以指向树中的任何节点或者指向空（null）。
//
// 请返回该树的 深拷贝 。
//
// 该树的输入/输出形式与普通二叉树相同，每个节点都用 [val, random_index] 表示：
//
//
// val：表示 Node.val 的整数
// random_index：随机指针指向的节点（在输入的树数组中）的下标；如果未指向任何节点，则为 null 。
//
//
// 该树以 Node 类的形式给出，而你需要以 NodeCopy 类的形式返回克隆得到的树。NodeCopy 类和Node 类定义一致。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [[1,null],null,[4,3],[7,0]]
//输出：[[1,null],null,[4,3],[7,0]]
//解释：初始二叉树为 [1,null,4,7] 。
//节点 1 的随机指针指向 null，所以表示为 [1, null] 。
//节点 4 的随机指针指向 7，所以表示为 [4, 3] 其中 3 是树数组中节点 7 对应的下标。
//节点 7 的随机指针指向 1，所以表示为 [7, 0] 其中 0 是树数组中节点 1 对应的下标。
//
//
// 示例 2：
//
//
//
//
//输入：root = [[1,4],null,[1,0],null,[1,5],[1,5]]
//输出：[[1,4],null,[1,0],null,[1,5],[1,5]]
//解释：节点的随机指针可以指向它自身。
//
//
// 示例 3：
//
//
//
//
//输入：root = [[1,6],[2,5],[3,4],[4,3],[5,2],[6,1],[7,0]]
//输出：[[1,6],[2,5],[3,4],[4,3],[5,2],[6,1],[7,0]]
//
//
//
//
// 提示：
//
//
// tree 中节点数目范围是 [0, 1000]
// 每个节点的值的范围是 [1, 10^6]
//
// Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 18 👎 0

// 2022-07-04 08:20:35
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

func copyRandomBinaryTree(root *Node) *NodeCopy {
	nodeCopies := make(map[*Node]*NodeCopy)

	var dfs func(*Node) *NodeCopy
	dfs = func(node *Node) *NodeCopy {
		if node == nil {
			return nil
		}
		if clone, ok := nodeCopies[node]; ok {
			return clone
		}
		duplicate := &NodeCopy{Val: node.Val}
		nodeCopies[node] = duplicate
		duplicate.Left = dfs(node.Left)
		duplicate.Right = dfs(node.Right)
		duplicate.Random = dfs(node.Random)

		return duplicate
	}

	return dfs(root)

}

//leetcode submit region end(Prohibit modification and deletion)
