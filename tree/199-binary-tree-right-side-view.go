package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		stack []*TreeNode = []*TreeNode{root}
		res   []int
		mid   []*TreeNode
	)
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1].Val)
		for _, e := range stack {
			if e.Left != nil {
				mid = append(mid, e.Left)
			}
			if e.Right != nil {
				mid = append(mid, e.Right)
			}
		}
		stack = mid
		mid = []*TreeNode{}
	}
	return res
}
