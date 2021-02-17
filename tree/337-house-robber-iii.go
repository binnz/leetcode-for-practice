package main

func rob(root *TreeNode) int {
	var (
		dfs func(*TreeNode) int
		res int
		mid map[*TreeNode]int = map[*TreeNode]int{}
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if e, ok := mid[node]; ok {
			return e
		}
		l, r := 0, 0
		if node.Left != nil {
			l = dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		if node.Right != nil {
			r = dfs(node.Right.Left) + dfs(node.Right.Right)
		}
		s1, s2 := 0, 0
		s1 = l + r + node.Val
		s2 = dfs(node.Right) + dfs(node.Left)
		res = max(s1, s2)
		return res
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
