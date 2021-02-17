package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	l := trimBST(root.Left, low, root.Val)
	r := trimBST(root.Right, root.Val, high)
	root.Left = l
	root.Right = r
	return root
}
