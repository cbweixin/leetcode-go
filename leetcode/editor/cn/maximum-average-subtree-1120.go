package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var dfs func(*TreeNode) (int, int)
	res := 0.0
	max := func(a, b float64) float64 {
		if a >= b {
			return a
		}
		return b
	}
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		sL, nL := dfs(node.Left)
		sR, nR := dfs(node.Right)
		res = max(res, float64(sL+sR+node.Val)/float64((nL+nR+1)))
		return sL + sR + node.Val, nL + nR + 1
	}
	dfs(root)

	return res

}
