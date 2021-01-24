package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var i, j int
	if nums[0] == val {
		i, j = 0, 1
	} else {
		i, j = 1, 1
	}
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
