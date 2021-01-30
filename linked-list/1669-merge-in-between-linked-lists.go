package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var cur, start, end *ListNode = list1, nil, nil
	for a > 0 || b+2 > 0 {
		a--
		b--
		if a == 0 {
			start = cur
		}
		if b == -2 {
			end = cur
		}
		cur = cur.Next
	}
	cur = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	start.Next = list2
	cur.Next = end
	return list1
}
