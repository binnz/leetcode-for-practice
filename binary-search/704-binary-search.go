package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 {
		return -1
	}
	if left == len(nums) {
		if nums[left-1] == target {
			return left - 1
		}
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}
