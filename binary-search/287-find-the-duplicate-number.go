package main

import "sort"

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	sort.Ints(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
