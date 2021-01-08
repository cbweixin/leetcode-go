package main

//Given the roots of two binary trees p and q, write a function to check if they
// are the same or not.
//
// Two binary trees are considered the same if they are structurally identical,
//and the nodes have the same value.
//
//
// Example 1:
//
//
//Input: p = [1,2,3], q = [1,2,3]
//Output: true
//
//
// Example 2:
//
//
//Input: p = [1,2], q = [1,null,2]
//Output: false
//
//
// Example 3:
//
//
//Input: p = [1,2,1], q = [1,1,2]
//Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in both trees is in the range [0, 100].
// -104 <= Node.val <= 104
//
// Related Topics Tree Depth-first Search
// 👍 2804 👎 76

// 2021-01-07 19:59:56

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	ch1, ch2 := make(chan int), make(chan int)

	go walk(p, ch1)
	go walk(q, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	_, ok := <-ch2
	if ok {
		return false
	}

	return true

}

func walk(n *TreeNode, ch chan int) {
	walkTree(n, ch)
	close(ch)
}

func walkTree(n *TreeNode, ch chan int) {
	if n != nil {
		ch <- n.Val
		walkTree(n.Left, ch)
		walkTree(n.Right, ch)
	} else {
		// send -1 if node is nil
		ch <- -1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
