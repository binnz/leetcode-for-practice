package main

func findLHS(nums []int) int {
	hm := map[int]int{}
	res := 0
	for _, e := range nums {
		hm[e]++
		if hm[e+1] > 0 {
			res = max(res, hm[e]+hm[e+1])
		}
		if hm[e-1] > 0 {
			res = max(res, hm[e]+hm[e-1])
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
