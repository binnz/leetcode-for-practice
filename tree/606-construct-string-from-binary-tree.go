package main

import "fmt"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	l := tree2str(t.Left)
	r := tree2str(t.Right)
	if l == "" && r == "" {
		return fmt.Sprintf("%d", t.Val)
	} else if r == "" {
		return fmt.Sprintf("%d(%s)", t.Val, l)
	} else {
		return fmt.Sprintf("%d(%s)(%s)", t.Val, l, r)
	}

}
