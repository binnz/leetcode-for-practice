package main

func findShortestSubArray(nums []int) int {
	numCount := map[int]int{}
	firstIdx := map[int]int{}
	lastIdx := map[int]int{}
	res := len(nums)
	for i, e := range nums {
		if _, ok := numCount[e]; ok {
			numCount[e]++
		} else {
			numCount[e]++
			firstIdx[e] = i
		}
		lastIdx[e] = i
	}
	maxNum := []int{}
	maxC := 0
	for k, v := range numCount {
		if v > maxC {
			maxC = v
			maxNum = []int{k}
		} else if v == maxC {
			maxNum = append(maxNum, k)
		}
	}
	for _, e := range maxNum {
		res = min(res, lastIdx[e]-firstIdx[e]+1)
	}
	return res

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
