package main

func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}
func minSteps1(n int) int {
	res := 0
	d := 2
	for n > 1 {
		for n%d == 0 {
			res += d
			n /= d
		}
		d++
	}
	return res
}
