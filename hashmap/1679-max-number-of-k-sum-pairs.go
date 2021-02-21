package main

import "sort"

func maxOperations(nums []int, k int) int {
	res := 0
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == k {
			res++
			i++
			j--
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}
	return res
}
