package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var (
		dfs func(*TreeNode, int)
		xf  *TreeNode
		yf  *TreeNode
		xd  int
		yd  int
	)
	if root.Val == x || root.Val == y {
		return false
	}
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if node.Left != nil {
			if node.Left.Val == x {
				xd = dep + 1
				xf = node
			}
			if node.Left.Val == y {
				yd = dep + 1
				yf = node
			}
			dfs(node.Left, dep+1)
		}
		if node.Right != nil {
			if node.Right.Val == x {
				xd = dep + 1
				xf = node
			}
			if node.Right.Val == y {
				yd = dep + 1
				yf = node
			}
			dfs(node.Right, dep+1)
		}
	}
	dfs(root, 0)
	if xf != yf && xd == yd {
		return true
	} else {
		return false
	}
}
