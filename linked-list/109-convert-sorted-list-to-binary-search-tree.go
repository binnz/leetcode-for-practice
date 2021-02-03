package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	fast, slow := head, head
	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	fast = slow.Next
	slow.Next = nil
	left := sortedListToBST(head)
	right := sortedListToBST(fast)
	root := &TreeNode{Val: slow.Val}
	root.Left = left
	root.Right = right
	return root

}
