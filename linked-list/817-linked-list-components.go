package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	res, start, valMap := 0, false, make(map[int]int)
	for _, e := range G {
		valMap[e] = 1
	}
	for head != nil {
		if _, ok := valMap[head.Val]; ok {
			start = true
		} else if start {
			res++
			start = false
		}
		head = head.Next

	}
	if start {
		res++
	}
	return res
}
