/**
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
**/
// Package main provides ...
package main

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums)
	var start int
	var end int

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == len(nums) {
		return []int{-1, -1}
	}
	if nums[left] == target {
		start = left
	} else {
		return []int{-1, -1}
	}
	left = 0
	right = len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}

		end = left - 1
	}
	return []int{start, end}
}
