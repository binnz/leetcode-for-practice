package main

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	cur := head
	if head == nil || head.Next == nil {
		return head
	}
	for cur != nil {
		cur = cur.Next
		length++
	}

	k = length - (k % length)
	if k == length {
		return head
	}
	cur = head
	for k > 1 {
		cur = cur.Next
		k--
	}
	start := cur.Next
	cur.Next = nil
	cur = start
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return start

}
