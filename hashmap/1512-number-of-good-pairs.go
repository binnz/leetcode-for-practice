package main

func numIdenticalPairs(nums []int) int {
	var idxToIdx map[int][]int = map[int][]int{}
	var numToIdx map[int][]int = map[int][]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if e, ok := numToIdx[nums[i]]; ok {
			idxToIdx[i] = e
			mid := []int{}
			mid = append(mid, e...)
			mid = append(mid, i)
			numToIdx[nums[i]] = mid
		} else {
			numToIdx[nums[i]] = []int{i}
		}
	}
	res := 0
	for i, _ := range nums {
		if v, ok := idxToIdx[i]; ok {
			res += len(v)
		}
	}
	return res
}
