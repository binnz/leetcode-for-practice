package main

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		stack []*TreeNode = []*TreeNode{root}
		mid   []*TreeNode
		res   []int
	)
	for len(stack) > 0 {
		var max int
		for i := 0; i < len(stack); i++ {
			e := stack[i]
			if i == 0 {
				max = e.Val
			}
			if e.Val > max {
				max = e.Val
			}
			if e.Left != nil {
				mid = append(mid, e.Left)
			}
			if e.Right != nil {
				mid = append(mid, e.Right)
			}
		}
		stack = mid
		res = append(res, max)
		mid = []*TreeNode{}
	}
	return res
}
