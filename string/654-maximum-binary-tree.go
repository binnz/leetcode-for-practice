package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if len(nums) == 0 {
		return nil
	}
	maxV := 0
	idx := 0
	for i, e := range nums {
		if e > maxV {
			maxV = e
			idx = i
		}
	}
	left := constructMaximumBinaryTree(nums[:idx])
	right := constructMaximumBinaryTree(nums[idx+1:])
	root := &TreeNode{Val: nums[idx], Left: left, Right: right}
	return root
}
