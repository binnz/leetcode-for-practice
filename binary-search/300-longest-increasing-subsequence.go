package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	res := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max2(res, dp[i])
	}
	return res

}

func lengthOfLIS1(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	res := 0
	for _, e := range nums {
		l, r := 0, res
		for l < r {
			mid := l + (r-l)>>1
			if dp[mid] >= e {
				r = mid
			} else {
				l = mid + 1
			}
		}
		dp[l] = e
		if l == res {
			res++
		}
	}
	return res
}
