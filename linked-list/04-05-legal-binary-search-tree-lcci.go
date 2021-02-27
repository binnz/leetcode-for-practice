package main

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode)
	flag := false
	pre := 0
	res := true
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if !flag {
			flag = true
			pre = node.Val
		} else {
			if node.Val <= pre {
				res = false
				return
			}
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}
