package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	var (
		dfs func(*TreeNode, int)
		res bool
	)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		if res {
			return
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if targetSum == sum {
				res = true
			}
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return res
}
