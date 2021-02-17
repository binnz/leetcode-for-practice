package main

func findBottomLeftValue(root *TreeNode) int {
	var (
		dfs   func(*TreeNode, int)
		res   int = root.Val
		level int
	)
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if dep > level {
			res = node.Val
			level = dep
		}
		dfs(node.Left, dep+1)
		dfs(node.Right, dep+1)
	}
	dfs(root, 0)
	return res

}
