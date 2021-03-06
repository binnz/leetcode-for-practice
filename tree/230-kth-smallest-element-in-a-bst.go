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

func kthSmallest1(root *TreeNode, k int) int {
	var reverse func(*TreeNode) func() int

	reverse = func(node *TreeNode) func() int {
		stack := []*TreeNode{root}
		n := []bool{false}
		return func() int {
			for len(stack) != 0 {
				curNode := stack[len(stack)-1]
				visited := n[len(n)-1]
				stack = stack[:len(stack)-1]
				n = n[:len(n)-1]
				if visited {
					return curNode.Val
				} else {
					if curNode.Right != nil {
						stack = append(stack, curNode.Right)
						n = append(n, false)
					}
					stack = append(stack, curNode)
					n = append(n, true)
					if curNode.Left != nil {
						stack = append(stack, curNode.Left)
						n = append(n, false)
					}
				}
			}
			return 0
		}
	}
	res := 0
	r := reverse(root)
	for i := 0; i < k; i++ {
		res = r()
	}
	return res
}
