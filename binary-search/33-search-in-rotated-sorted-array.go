package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] > nums[len(nums)-1] {
			l = m + 1
		} else {
			r = m
		}
	}
	var s, e = 0, 0
	if nums[0] < nums[len(nums)-1] {
		s, e = 0, len(nums)
	} else if target >= nums[0] {
		s, e = 0, l-1
	} else {
		s, e = l, len(nums)
	}
	l, r = s, e
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == s || l == e {
		return -1
	}
	if nums[l] == target {
		return l
	}
	return -1
}
