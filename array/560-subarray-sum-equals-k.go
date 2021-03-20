package main

func subarraySum(nums []int, k int) int {
	hm := map[int]int{}
	sum := 0
	res := 0
	hm[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := hm[sum-k]; ok {
			res += v
		}
		hm[sum]++
	}
	return res
}
