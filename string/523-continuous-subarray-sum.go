package main

func checkSubarraySum(nums []int, k int) bool {
	mod := map[int]int{}
	sum := map[int]int{}
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && nums[i+1] == 0 {
			return true
		}
	}
	if k == 0 {
		return false
	}
	if k < 0 {
		k = -k
	}
	c := 0
	mod[0] = -1
	for i := 0; i < len(nums); i++ {
		c += nums[i]
		sum[i] = c
		m := sum[i] % k
		if v, ok := mod[m]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			mod[m] = i
		}
	}
	return false
}
