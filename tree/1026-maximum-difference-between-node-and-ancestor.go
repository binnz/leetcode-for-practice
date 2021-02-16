package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxAncestorDiff(root *TreeNode) int {
	var (
		dfs func(*TreeNode, int, int)
		res int
	)
	dfs = func(node *TreeNode, max, min int) {
		if node == nil {
			return
		}
		if node.Val < min {
			min = node.Val
		}
		if node.Val > max {
			max = node.Val
		}
		if max-min > res {
			res = max - min
		}
		if node.Left != nil {
			dfs(node.Left, max, min)
		}
		if node.Right != nil {
			dfs(node.Right, max, min)
		}
	}
	dfs(root, root.Val, root.Val)
	return res
}
