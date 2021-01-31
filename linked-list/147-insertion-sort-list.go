package main

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	start := dummy
	for head != nil {
		mid := head.Next
		cur, prev := start, start
		for cur != dummy && head.Val < cur.Val {
			prev = cur
			cur = cur.Next
		}
		if cur == prev {
			head.Next = cur
			start = head
		} else {
			head.Next = cur
			prev.Next = head
		}
		head = mid
	}
	var res *ListNode
	for start != nil {
		mid := start.Next
		start.Next = res
		res = start
		start = mid
	}
	return res.Next
}
