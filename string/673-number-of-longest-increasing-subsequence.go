package main

func findNumberOfLIS(nums []int) int {
	maxL := 1
	res := 1
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max2(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxL {
			maxL = dp[i]
			res = 1
		} else if dp[i] == maxL {
			res++
		}
	}
	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
