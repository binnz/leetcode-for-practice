package main

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(int, int) *TreeNode

	dfs = func(left int, right int) *TreeNode {
		if left == right {
			return &TreeNode{Val: nums[left]}
		}
		if left > right {
			return nil
		}
		mid := left + (right-left)>>1
		root := &TreeNode{Val: nums[mid]}
		l := dfs(left, mid-1)
		r := dfs(mid+1, right)
		root.Left = l
		root.Right = r
		return root
	}
	return dfs(0, len(nums)-1)
}
