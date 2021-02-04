package main

func partition(head *ListNode, x int) *ListNode {
	first := &ListNode{Val: 0}
	second := &ListNode{Val: 0}
	curLess := first
	curGreat := second
	for head != nil {
		var mid *ListNode
		if head.Val < x {
			mid = head.Next
			curLess.Next = head
			curLess = head
		} else {
			mid = head.Next
			curGreat.Next = head
			curGreat = head
		}
		head = mid
	}
	curGreat.Next = nil
	curLess.Next = second.Next
	return first.Next
}
