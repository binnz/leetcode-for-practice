package main

func averageOfLevels(root *TreeNode) []float64 {
	var (
		stack []*TreeNode = []*TreeNode{root}
		mid   []*TreeNode = []*TreeNode{}
		res   []float64
	)
	for len(stack) != 0 {
		c := len(stack)
		var sum float64 = 0
		for i := 0; i < c; i++ {
			e := stack[i]
			if e.Left != nil {
				mid = append(mid, e.Left)
			}
			if e.Right != nil {
				mid = append(mid, e.Right)
			}
			sum += float64(e.Val)
		}
		res = append(res, sum/float64(c))
		stack = mid
		mid = []*TreeNode{}
	}
	return res
}
