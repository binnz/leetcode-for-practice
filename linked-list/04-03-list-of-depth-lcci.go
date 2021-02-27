package main

func listOfDepth(tree *TreeNode) []*ListNode {
	res := []*ListNode{}
	stack := []*TreeNode{tree}
	for len(stack) != 0 {
		mid := []*TreeNode{}
		dummy := &ListNode{}
		pre := dummy
		for i := 0; i < len(stack); i++ {
			node := stack[i]
			if node.Left != nil {
				mid = append(mid, node.Left)
			}
			if node.Right != nil {
				mid = append(mid, node.Right)
			}
			dummy.Next = &ListNode{Val: node.Val}
			dummy = dummy.Next
		}
		res = append(res, pre.Next)

		stack = mid
	}
	return res
}
