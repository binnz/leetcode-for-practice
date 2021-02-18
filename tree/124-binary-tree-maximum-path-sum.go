package main

func maxPathSum(root *TreeNode) int {
	var (
		res int = -100000000
		dfs func(*TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(dfs(node.Left), 0)
		r := max(dfs(node.Right), 0)
		res = max(res, l+r+node.Val)
		return max(l, r) + node.Val
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
