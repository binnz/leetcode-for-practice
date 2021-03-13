package main

func findLengthOfLCIS(nums []int) int {
	res := 1
	i, j := 0, 1
	if len(nums) < 2 {
		return len(nums)
	}
	for j < len(nums) {
		if nums[j] > nums[j-1] {
			res = max(res, j-i+1)
			j++
		} else {
			i = j
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
