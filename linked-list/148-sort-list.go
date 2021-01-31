package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	head = sortList(head)
	slow = sortList(slow)
	res := merge(head, slow)
	return res
}

func merge(a, b *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	res := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			dummy.Next = a
			a = a.Next
		} else {
			dummy.Next = b
			b = b.Next
		}
		dummy = dummy.Next
	}
	if a != nil {
		dummy.Next = a
	}
	if b != nil {
		dummy.Next = b
	}
	return res.Next
}
