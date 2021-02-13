package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	c := 1 << len(nums)
	for i := 0; i < c; i++ {
		mid := []int{}
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				mid = append(mid, nums[j])
			}
		}
		res = append(res, mid)
	}
	return res

}
