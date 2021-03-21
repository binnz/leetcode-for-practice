package main

// func findLongestChain(pairs [][]int) int {
// 	dp := make([]int, len(pairs)+1)
// 	res := 0
// 	arr := Pairs{}
// 	arr = pairs
// 	sort.Sort(arr)
// 	for i, _ := range pairs {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < len(arr); i++ {
// 		for j := 0; j < i; j++ {
// 			if arr[i][0] > arr[j][1] {
// 				dp[i] = max4(dp[i], dp[j]+1)
// 				res = max4(res, dp[i])
// 			}
// 		}
// 	}
// 	return res
// }
// func max4(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type Pairs [][]int

// func (p Pairs) Len() int {
// 	return len(p)
// }

// func (p Pairs) Less(i, j int) bool {
// 	if p[i][0] == p[j][0] {
// 		return p[i][1] < p[j][1]
// 	}
// 	return p[i][0] < p[j][0]
// }
// func (p Pairs) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }
