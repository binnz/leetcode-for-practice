package main

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		sum    int
		stack  []*TreeNode = []*TreeNode{root}
		stackB []bool      = []bool{false}
	)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		curState := stackB[len(stackB)-1]
		stack = stack[:len(stack)-1]
		stackB = stackB[:len(stackB)-1]
		if curState {
			mid := curNode.Val
			curNode.Val += sum
			sum += mid
		} else {
			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
				stackB = append(stackB, false)
			}
			stack = append(stack, curNode)
			stackB = append(stackB, true)
			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
				stackB = append(stackB, false)
			}
		}
	}
	return root
}
