package main

func findFrequentTreeSum(root *TreeNode) []int {
	var (
		dfs func(*TreeNode) int
		res []int
		mid map[int]int = map[int]int{}
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		v := l + r + node.Val
		if _, ok := mid[v]; ok {
			mid[v]++
		} else {
			mid[v] = 1
		}
		return v
	}
	dfs(root)
	var s int
	for k, v := range mid {
		if v > s {
			res = []int{k}
			s = v
		} else if v == s {
			res = append(res, k)
		}
	}
	return res
}
