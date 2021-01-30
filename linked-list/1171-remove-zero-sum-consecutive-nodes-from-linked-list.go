package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	sumToNode := make(map[int]*ListNode)
	dummy := &ListNode{Val: 0, Next: head}
	cur, sums := dummy, 0
	for cur != nil {
		sums += cur.Val
		sumToNode[sums] = cur
		cur = cur.Next
	}
	cur, sums = dummy, 0
	for cur != nil {
		sums += cur.Val
		cur.Next = sumToNode[sums].Next
		cur = cur.Next
	}
	return dummy.Next

}
