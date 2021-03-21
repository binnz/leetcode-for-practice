package main

func findMaxLength(nums []int) int {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	hm := map[int]int{}
	res := 0
	sum := 0
	hm[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := hm[sum]; ok {
			res = max(res, i-v)
		}
		if _, ok := hm[sum]; !ok {
			hm[sum] = i
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
