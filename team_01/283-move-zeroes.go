package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var i, j int
	if nums[0] == 0 {
		i, j = 0, 1
	} else {
		i, j = 1, 1
	}
	for j < len(nums) {
		if nums[j] != 0 {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
		j++
	}
}
