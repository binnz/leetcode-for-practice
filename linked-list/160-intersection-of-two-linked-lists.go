package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	first, second := headA, headB
	for first != second {
		if first != nil {
			first = first.Next
		} else {
			first = headB
		}
		if second != nil {
			second = second.Next
		} else {
			second = headA
		}

	}
	return first
}
