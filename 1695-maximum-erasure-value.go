package main

func maximumUniqueSubarray(nums []int) int {
	addSum := make(map[int]int)
	addSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		addSum[i] = nums[i] + addSum[i-1]
	}
	numSet := make(map[int]int)
	prevSameEle := make(map[int]int)
	for idx, e := range nums {
		if i, ok := numSet[e]; ok {
			prevSameEle[idx] = i
			numSet[e] = idx
		} else {
			prevSameEle[idx] = -1
			numSet[e] = idx
		}
	}
	i, j, res := 0, 1, nums[0]
	for j < len(nums) {
		if idx, _ := prevSameEle[j]; idx >= i {
			i = idx + 1
		}
		mid := addSum[j] - addSum[i] + nums[i]
		if mid > res {
			res = mid
		}
		j++
	}
	return res
}
