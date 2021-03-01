package main

import "math"

func smallestDivisor(nums []int, threshold int) int {
	left, right := 0, 1000000
	for left < right {
		mid := left + (right-left)>>1
		if calSum(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 {
		return 1
	}
	return left
}

func calSum(nums []int, divisor int) int {
	res := 0.0
	for _, e := range nums {
		res += math.Ceil(float64(e) / float64(divisor))
	}
	return int(res)
}
