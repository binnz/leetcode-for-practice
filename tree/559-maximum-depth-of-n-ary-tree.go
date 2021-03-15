package main

func maxDepth(root *Node) int {
	res := 0
	if root == nil {
		return 0
	}
	var dfs func(*Node, int)
	dfs = func(node *Node, d int) {

		d++
		if len(node.Children) == 0 {
			res = max(res, d)
			return
		}
		for _, e := range node.Children {
			dfs(e, d)
		}

	}
	dfs(root, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
