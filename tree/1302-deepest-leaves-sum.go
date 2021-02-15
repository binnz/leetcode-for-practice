package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	maxDeep, res := 0, 0
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	deep := []int{0}
	for len(stack) != 0 {
		curNode := stack[len(stack)-1]
		curDeep := deep[len(deep)-1]
		stack = stack[:len(stack)-1]
		deep = deep[:len(deep)-1]
		if curDeep == maxDeep {
			res += curNode.Val
		} else if curDeep > maxDeep {
			maxDeep = curDeep
			res = curNode.Val
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
			deep = append(deep, curDeep+1)
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
			deep = append(deep, curDeep+1)
		}
	}
	return res
}
