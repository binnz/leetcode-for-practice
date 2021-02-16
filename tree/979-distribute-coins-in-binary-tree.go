package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	var (
		dfs func(*TreeNode) int
		res int
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		res += (abs(l) + abs(r))
		return node.Val + l + r - 1
	}
	dfs(root)
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
