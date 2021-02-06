package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	for second != nil {
		if first.Val != second.Val {
			first = first.Next
			second = second.Next
		} else {
			second = second.Next
			first.Next = second
		}
	}
	return head
}
