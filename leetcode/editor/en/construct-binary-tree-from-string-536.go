package main

import "strconv"

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

func str2tree(s string) *TreeNode {
	l, st, i := len(s), []TreeNode{}, 0

	is_digit := func(x uint8) bool {
		return x >= '0' && x <= '9'
	}

	for i < l {
		if s[i] == ')' {
			// pop
			st = st[:len(st)-1]
		} else if is_digit(s[i]) || s[i] == '-' {
			k := i + 1
			for is_digit(s[k]) {
				k++
			}

			n, _ := strconv.Atoi(s[i:k])
			node := &TreeNode{n, nil, nil}
			if len(st) > 0 {
				if st[len(st)-1].Left == nil {
					st[len(st)-1].Left = node
				} else {
					st[len(st)-1].Right = node
				}
			}
			st = append(st, *node)
			i = k - 1
		}
		i++
	}

	return &st[0]

}

func main() {

}
