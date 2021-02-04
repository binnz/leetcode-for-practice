package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	for m > 1 {
		prev = prev.Next
		m--
		n--
	}
	current := prev.Next
	var pre *ListNode
	for n > 0 {
		mid := current.Next
		current.Next = pre
		pre = current
		current = mid
		n--
	}
	prev.Next.Next = current
	prev.Next = pre
	return dummy.Next
}
