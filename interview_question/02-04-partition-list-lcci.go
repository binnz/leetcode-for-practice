package main

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	dummy1 := &ListNode{}
	cur1 := dummy1
	for head != nil {
		if head.Val < x {
			next := head.Next
			cur.Next = head
			cur = cur.Next
			head.Next = nil
			head = next
		} else {
			next := head.Next
			cur1.Next = head
			cur1 = cur1.Next
			head.Next = nil
			head = next
		}
	}
	cur.Next = dummy1.Next
	return dummy.Next
}
