package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.NExt == nil {
		return head
	}
	var dummy *ListNode
	for head != nil {
		mid := head.Next
		head.Next = dummy
		dummy = head
		head = mid
	}
	return dummy
}
