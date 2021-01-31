package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even = head.Next
	evenStart := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenStart
	return head
}
