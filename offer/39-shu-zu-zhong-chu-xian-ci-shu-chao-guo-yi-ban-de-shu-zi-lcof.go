package main

func majorityElement(nums []int) int {
	res := 0
	c := 0
	for i := 0; i < len(nums); i++ {
		if c == 0 {
			res = nums[i]
			c++
		} else {
			if nums[i] == res {
				c++
			} else {
				c--
			}
		}
	}
	return res
}
