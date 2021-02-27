package main

func removeDuplicateNodes(head *ListNode) *ListNode {
	res := head
	if head == nil {
		return head
	}
	for head.Next != nil {
		node := head
		cur := head
		for cur.Next != nil {
			if cur.Next.Val == node.Val {
				cur.Next = cur.Next.Next
				if cur.Next != nil {
					continue
				}
			}
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}

		}
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}

	}
	return res
}
