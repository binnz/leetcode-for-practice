package main

func maximumSum(arr []int) int {
	if len(arr) < 2 {
		return arr[0]
	}
	l := len(arr)
	dp := make([]int, l)
	dp2 := make([]int, l)
	res := -l * 10001
	dp[0] = arr[0]
	dp2[0] = -2 * 20001
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+arr[i], arr[i])
		dp2[i] = max(dp2[i-1]+arr[i], dp[i-1])
		res = max(res, max(dp[i], dp2[i]))
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
