package main

func findDuplicates(nums []int) []int {
	res := []int{}
	for _, e := range nums {
		if nums[abs(e)-1] > 0 {
			nums[abs(e)-1] *= -1
		} else {
			res = append(res, abs(e))
		}
	}
	return res
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
