package main

func flatten(root *TreeNode) {
	var dfs func(*TreeNode) (*TreeNode, *TreeNode)
	dfs = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root == nil {
			return nil, nil
		}
		if root.Left == nil && root.Right == nil {
			return root, root
		}
		mid := root.Right
		left, end := dfs(root.Left)
		right, rend := dfs(mid)
		root.Left = nil
		if left != nil {
			root.Right = left
			end.Right = right
			if right != nil {
				return root, rend
			} else {
				return root, end
			}
		} else {
			root.Right = right
			if right != nil {
				return root, rend
			} else {
				return root, root
			}
		}
	}
	dfs(root)
}
