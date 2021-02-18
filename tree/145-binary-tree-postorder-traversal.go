package main

func postorderTraversal(root *TreeNode) []int {
	var (
		dfs func(*TreeNode)
		res []int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}
