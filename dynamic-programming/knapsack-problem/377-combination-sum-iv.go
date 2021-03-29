package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, e := range nums {
			if i >= e {
				dp[i] += dp[i-e]
			}
		}
	}
	return dp[target]
}
