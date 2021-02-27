package main

func kthToLast(head *ListNode, k int) int {
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	for l-k > 0 {
		head = head.Next
		k++
	}
	return head.Val
}
