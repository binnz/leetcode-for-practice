package main

func pathSum(root *TreeNode, sum int) int {

	var (
		dfs  func(*TreeNode)
		dfss func(*TreeNode, int, int)
		res  int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfss(node, 0, sum)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfss = func(node *TreeNode, sum int, total int) {
		if node == nil {
			return
		}

		sum += node.Val
		if sum == total {
			res++
		}
		dfss(node.Left, sum, total)
		dfss(node.Right, sum, total)
	}
	dfs(root)
	return res
}
