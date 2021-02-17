package main

func sumOfLeftLeaves(root *TreeNode) int {
	var (
		dfs func(*TreeNode)
		res int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
