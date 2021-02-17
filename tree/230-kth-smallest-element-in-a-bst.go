package main

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var (
		stack []*TreeNode = []*TreeNode{root}
		visit []bool      = []bool{false}
	)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		visited := visit[len(visit)-1]
		stack = stack[:len(stack)-1]
		visit = visit[:len(visit)-1]
		if visited {
			if k == 1 {
				return node.Val
			}
			k--
		} else {
			if node.Right != nil {
				stack = append(stack, node.Right)
				visit = append(visit, false)
			}
			stack = append(stack, node)
			visit = append(visit, true)
			if node.Left != nil {
				stack = append(stack, node.Left)
				visit = append(visit, false)
			}
		}
	}
	return 0

}
