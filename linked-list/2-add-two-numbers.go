package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	h := 0
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val + h
		h = v / 10
		vv := v % 10
		dummy.Next = &ListNode{Val: vv}
		dummy = dummy.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	var t *ListNode
	if l1 != nil {
		t = l1
	} else if l2 != nil {
		t = l2
	} else {
		if h != 0 {
			dummy.Next = &ListNode{Val: h}
		}
		return res.Next
	}
	dummy.Next = t
	dummy = dummy.Next
	for h != 0 {
		fmt.Println(h, dummy.Val)
		v := dummy.Val + h
		h = v / 10
		vv := v % 10
		dummy.Val = vv
		if dummy.Next == nil {
			break
		}
		dummy = dummy.Next
	}
	if h != 0 {
		dummy.Next = &ListNode{Val: h}
	}
	return res.Next
}
