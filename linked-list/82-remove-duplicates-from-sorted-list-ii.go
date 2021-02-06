package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 1, Next: head}
	cur, pre := head, dummy
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			pre.Next = cur
			pre = cur
			cur = cur.Next
		} else {
			val := cur.Val
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			pre.Next = cur
			pre = cur
		}
	}
	return dummy.Next

}
