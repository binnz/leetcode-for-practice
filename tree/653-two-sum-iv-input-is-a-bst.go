package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	var (
		hm  map[int]struct{} = map[int]struct{}{}
		dfs func(*TreeNode)
		res bool
	)
	dfs = func(node *TreeNode) {
		if node == nil || res {
			return
		}
		dfs(node.Left)
		if _, ok := hm[k-node.Val]; ok {
			res = true
		} else {
			hm[node.Val] = struct{}{}
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
