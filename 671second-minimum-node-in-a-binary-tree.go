package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	var dfs func(*TreeNode, int) int

	dfs = func(node *TreeNode, v int) int {
		if node == nil {
			return -1
		}
		if node.Val > v {
			return node.Val
		}
		l := dfs(node.Left, v)
		r := dfs(node.Right, v)
		if l > v && r > v {
			if l > r {
				return r
			} else {
				return l
			}
		} else {
			if l > r {
				return l
			} else {
				return r
			}
		}
	}
	return dfs(root, root.Val)

}
