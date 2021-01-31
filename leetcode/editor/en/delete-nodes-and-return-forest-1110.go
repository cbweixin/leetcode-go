package main

//Given the root of a binary tree, each node in the tree has a distinct value.
//
// After deleting all nodes with a value in to_delete, we are left with a forest
// (a disjoint union of trees).
//
// Return the roots of the trees in the remaining forest. You may return the res
//ult in any order.
//
//
// Example 1:
//
//
//
//
//Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
//Output: [[1,2,null,4],[6],[7]]
//
//
//
// Constraints:
//
//
// The number of nodes in the given tree is at most 1000.
// Each node has a distinct value between 1 and 1000.
// to_delete.length <= 1000
// to_delete contains distinct values between 1 and 1000.
// Related Topics Tree Depth-first Search
// 👍 1572 👎 52

// 2021-01-28 09:45:49

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func dfs(node *TreeNode, parent *TreeNode, del_lookup map[int]bool, forest *[]*TreeNode) {
	if node == nil {
		return
	}
	if _, ok := del_lookup[node.Val]; ok {
		dfs(node.Left, nil, del_lookup, forest)
		dfs(node.Right, nil, del_lookup, forest)
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}

	} else {
		if parent == nil {
			*forest = append(*forest, node)
		}
		dfs(node.Left, node, del_lookup, forest)
		dfs(node.Right, node, del_lookup, forest)

	}

}
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	forest := make([]*TreeNode, 0)
	delete_lookup := make(map[int]bool)
	for _, v := range to_delete {
		delete_lookup[v] = true
	}

	dfs(root, nil, delete_lookup, &forest)
	return forest

}

//leetcode submit region end(Prohibit modification and deletion)
