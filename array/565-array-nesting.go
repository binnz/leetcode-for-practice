package main

func arrayNesting(nums []int) int {
	res := 0
	for i, e := range nums {
		if e == -1 {
			continue
		}
		n := i
		c := 0
		for n >= 0 && n < len(nums) && nums[n] != -1 {
			c++
			tmp := nums[n]
			nums[n] = -1
			n = tmp
		}
		res = max(res, c)
	}
	return res
}
