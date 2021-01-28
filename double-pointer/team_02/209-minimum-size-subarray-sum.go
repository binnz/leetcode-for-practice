package main

func minSubArrayLen(s int, nums []int) int {
	left, right, sums, res := 0, 0, 0, len(nums)+1
	for right < len(nums) {
		sums += nums[right]
		right++
		for sums >= s {
			res = min(res, right-left)
			sums -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
