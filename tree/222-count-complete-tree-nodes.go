package main

func countNodes(root *TreeNode) int {
	var (
		dfs func(*TreeNode)
		res int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
