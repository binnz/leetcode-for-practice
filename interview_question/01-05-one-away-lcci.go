package main

func oneEditAway(first string, second string) bool {
	row := len(first)
	col := len(second)
	if row == 0 {
		return col < 2
	}
	if col == 0 {
		return row < 2
	}
	dp := make([][]int, row+1)
	for k, _ := range dp {
		dp[k] = make([]int, col+1)
	}
	for i := 1; i < col+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < row+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[row][col] <= 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func oneEditAway2(first string, second string) bool {
	m, n := len(first), len(second)
	dif, max := m-n, m
	if n > m {
		dif = n - m
		max = n
		first, second = second, first
	}
	if dif > 1 {
		return false
	}
	c := 0
	if m == n {
		for i := 0; i < m; i++ {
			if first[i] != second[i] {
				c++
			}
			if c > 1 {
				return false
			}
		}
		return true
	} else {
		res := false
		for i := 0; i < max; i++ {
			if i == 0 {
				res = first[1:] == second
			} else if i == max-1 {
				res = first[:max-1] == second
			} else {
				res = first[:i] == second[:i] && first[i+1:] == second[i:]
			}
			if res {
				return res
			}
		}
		return res
	}
}
