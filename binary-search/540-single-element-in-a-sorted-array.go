package main

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	if len(nums) == 1 {
		return nums[0]
	}
	for l < r {
		mid := l + (r-l)>>1
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return nums[l]
}
