package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cur, length := head, 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	length /= 2
	for length > 0 {
		head = head.Next
		length--
	}
	return head

}
