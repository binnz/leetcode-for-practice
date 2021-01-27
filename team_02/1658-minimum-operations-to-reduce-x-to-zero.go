package main

func minOperations(nums []int, x int) int {
	total := 0
	for _, e := range nums {
		total += e
	}
	total -= x
	if total < 0 {
		return -1
	}
	if total == 0 {
		return len(nums)
	}
	left, right, sums, res := 0, 0, 0, -1
	for right < len(nums) {
		sums += nums[right]
		right++
		for sums >= total {
			if sums == total {
				res = max(res, right-left)
			}
			sums -= nums[left]
			left++
		}
	}
	if res == -1 {
		return res
	}
	return len(nums) - res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
