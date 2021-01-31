package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2, nums []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	pre := 0
	for len(stack1) > 0 || len(stack2) > 0 {
		mid, first, second := 0, 0, 0
		if len(stack1) > 0 && len(stack2) > 0 {
			first = stack1[len(stack1)-1]
			second = stack2[len(stack2)-1]
			stack1 = stack1[:len(stack1)-1]
			stack2 = stack2[:len(stack2)-1]
		} else if len(stack1) > 0 {
			first = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		} else {
			second = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		mid = first + second + pre
		pre = mid / 10
		mid = mid % 10

		nums = append([]int{mid}, nums...)
	}
	if pre == 1 {
		nums = append([]int{pre}, nums...)
	}
	dummy := &ListNode{Val: 0, Next: nil}
	res := dummy
	for _, e := range nums {

		node := &ListNode{Val: e, Next: nil}
		dummy.Next = node
		dummy = node
	}

	return res.Next
}
