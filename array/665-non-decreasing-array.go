package main

func checkPossibility(nums []int) bool {
	c := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			c++
			if i > 1 {
				if nums[i] >= nums[i-2] {
					nums[i-1] = nums[i-2]
				} else {
					nums[i] = nums[i-1]
				}
			}
		}
		if c == 2 {
			return false
		}
	}
	return true
}
