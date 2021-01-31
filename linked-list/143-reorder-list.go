package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	intToNodeMap, i := make(map[int]*ListNode), 0
	for head != nil {
		mid := head.Next
		head.Next = nil
		intToNodeMap[i] = head
		head = mid
		i++
	}
	left, right := 0, i-1
	dummy := &ListNode{Val: 0, Next: nil}
	res := dummy
	for left < right {
		one := intToNodeMap[left]
		two := intToNodeMap[right]
		dummy.Next = one
		one.Next = two
		dummy = two
		left++
		right--
	}
	if left == right {
		dummy.Next = intToNodeMap[left]
	}
	head = res.Next
}
