// Package main provides ...
/**
 https://leetcode-cn.com/problems/wiggle-subsequence/
**/
package main

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	res := 1
	pre_sign := nums[1] - nums[0]
	if pre_sign != 0 {
		res++
	}

	for i := 2; i < len(nums); i++ {
		sign := nums[i] - nums[i-1]
		if sign > 0 && pre_sign <= 0 || sign < 0 && pre_sign >= 0 {
			res++
			pre_sign = sign
		}
	}
	return res
}
