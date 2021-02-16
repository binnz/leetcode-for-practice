package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }

同538题 https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
*/
func bstToGst(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	var sum int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		mid := node.Val
		node.Val += sum
		sum += mid
		dfs(node.Left)
		return
	}
	dfs(root)
	return root
}

func bstToGst2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		sum    int
		stack  []*TreeNode = []*TreeNode{root}
		stackB []bool      = []bool{false}
	)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		curState := stackB[len(stackB)-1]
		stack = stack[:len(stack)-1]
		stackB = stackB[:len(stackB)-1]
		if curState {
			mid := curNode.Val
			curNode.Val += sum
			sum += mid
		} else {
			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
				stackB = append(stackB, false)
			}
			stack = append(stack, curNode)
			stackB = append(stackB, true)
			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
				stackB = append(stackB, false)
			}
		}
	}
	return root
}
