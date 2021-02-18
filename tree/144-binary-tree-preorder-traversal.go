package main

func preorderTraversal(root *TreeNode) []int {
	var (
		dfs func(*TreeNode)
		res []int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
