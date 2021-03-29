package main

func nextPermutation(nums []int) {
	l := len(nums)
	i := l - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i == 0 {
		for j := 0; j < l/2; j++ {
			nums[j], nums[l-j-1] = nums[l-j-1], nums[j]
		}
		return
	}
	target := i - 1

	for j := i; j < l; j++ {
		for k := j + 1; k < l; k++ {
			if nums[k] < nums[j] {
				nums[k], nums[j] = nums[j], nums[k]
			}
		}
	}
	j := i
	for nums[j] <= nums[target] {
		j++
	}
	nums[target], nums[j] = nums[j], nums[target]
	for j := i; j < l; j++ {
		for k := j + 1; k < l; k++ {
			if nums[k] < nums[j] {
				nums[k], nums[j] = nums[j], nums[k]
			}
		}
	}
	return
}
