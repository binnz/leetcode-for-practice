package main

// 输入：[2,7,4,3,5]
// 输出：[7,0,5,5,0]

func nextLargerNodes(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	stack, res := make([]int, len(nums)), make([]int, len(nums))
	for idx, e := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < e {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[i] = e
		}
		stack = append(stack, idx)
	}
	return res
}
