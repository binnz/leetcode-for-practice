package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		dfs func(*TreeNode, []int, int)
		res [][]int
	)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if targetSum == sum {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}
	dfs(root, []int{}, 0)
	return res
}
