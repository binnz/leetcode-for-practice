package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
