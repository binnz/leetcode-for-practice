package main

func findMin(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > nums[len(nums)-1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
