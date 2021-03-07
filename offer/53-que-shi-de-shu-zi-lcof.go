package main

func missingNumber(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] <= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
