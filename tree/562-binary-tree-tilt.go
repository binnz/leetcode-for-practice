package main

func findTilt(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var res int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		res += abs(l, r)
		return l + r + node.Val
	}

	dfs(root)
	return res

}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
