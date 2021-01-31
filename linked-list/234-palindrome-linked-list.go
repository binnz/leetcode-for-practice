package main

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	length, cur := 0, head
	for cur != nil {
		length++
		cur = cur.Next
	}
	sign := length % 2
	length /= 2
	cur = head
	var second *ListNode
	for length > 1 {
		cur = cur.Next
		length--
	}
	if sign == 1 {
		second = cur.Next.Next
		cur.Next.Next = nil
	} else {
		second = cur.Next
		cur.Next = nil
	}
	var dummy *ListNode
	for second != nil {
		mid := second.Next
		second.Next = dummy
		dummy = second
		second = mid
	}
	for head != nil && dummy != nil {
		if head.Val != dummy.Val {
			return false
		}
		head = head.Next
		dummy = dummy.Next
	}
	return true
}
