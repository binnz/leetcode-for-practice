package main

func minSubArrayLen(s int, nums []int) int {
	left, right, sums, res := 0, 0, 0, len(nums)+1
	for right < len(nums) {
		sums += nums[right]
		right++
		for sums >= s {
			res = min(res, right-left)
			sums -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen1(target int, nums []int) int {
	left, right := 0, len(nums)+1
	for left < right {
		mid := left + (right-left)>>1
		var ok bool
		sum := 0
		for i, e := range nums {
			if i < mid {
				sum += e
			} else {
				if sum >= target {
					ok = true
					break
				} else {
					sum -= nums[i-mid]
					sum += e
				}
			}
		}
		if sum >= target {
			ok = true
		}
		if !ok {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == len(nums)+1 {
		return 0
	}
	return left
}
