package main

func longestUnivaluePath(root *TreeNode) int {
	var dfs func(*TreeNode) int
	res := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftIncludeLen := dfs(node.Left)
		rightIncludeLen := dfs(node.Right)
		l, r := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			l = leftIncludeLen + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			r = rightIncludeLen + 1
		}
		res = max(res, l+r)
		return max(l, r)
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
