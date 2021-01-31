package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	cur, length, team, res := root, 0, make([]int, k), make([]*ListNode, k)
	for cur != nil {
		length++
		cur = cur.Next
	}
	i := 0
	for length > 0 {
		team[i]++
		length--
		if i < k-1 {
			i++
		} else {
			i = 0
		}
	}
	cur = root
	for idx, e := range team {
		if e == 0 {
			res[idx] = nil
		} else {
			for e > 1 {
				cur = cur.Next
				e--
			}
			mid := cur.Next
			cur.Next = nil
			res[idx] = root
			root = mid
			cur = root
		}
	}
	return res
}
