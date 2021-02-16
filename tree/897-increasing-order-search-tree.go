package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var (
		dfs func(*TreeNode)
		cur *TreeNode = &TreeNode{}
	)
	res := cur
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return res.Right
}
