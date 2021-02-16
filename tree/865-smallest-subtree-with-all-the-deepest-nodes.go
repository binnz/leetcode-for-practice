package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l == r {
		return root
	} else if l > r {
		return subtreeWithAllDeepest(root.Left)
	} else {
		return subtreeWithAllDeepest(root.Right)
	}
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
