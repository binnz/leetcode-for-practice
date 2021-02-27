package main

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(dfs(node.Left), dfs(node.Right)) + 1
	}
	if root == nil {
		return true
	}
	if abs(dfs(root.Left), dfs(root.Right)) >= 2 {
		return false
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
