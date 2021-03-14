package main

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	hm := make(map[string]int)
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		s := fmt.Sprintf("%d,%s,%s", node.Val, dfs(node.Left), dfs(node.Right))
		if hm[s] == 1 {
			res = append(res, node)
		}
		hm[s]++
		return s
	}
	dfs(root)
	return res
}
