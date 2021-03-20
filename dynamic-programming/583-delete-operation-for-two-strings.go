package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	l1 := len(word1)
	l2 := len(word2)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i < l1; i++ {
		dp[i+1][0] = i + 1
	}
	for i := 0; i < l2; i++ {
		dp[0][i+1] = i + 1
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[l1][l2]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
