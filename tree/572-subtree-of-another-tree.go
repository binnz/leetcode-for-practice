package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	m1 := isSubtree(s.Left, t) || isSubtree(s.Right, t)
	m2 := isSameTree(s, t)
	return m1 || m2

}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	} else {
		l := isSameTree(s.Left, t.Left)
		r := isSameTree(s.Right, t.Right)
		if l && r {
			return true
		} else {
			return false
		}
	}
}
