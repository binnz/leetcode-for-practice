package main

func sumOfUnique(nums []int) int {
	var (
		res int
		hm  map[int]int = map[int]int{}
	)
	for _, e := range nums {
		hm[e]++
	}
	for k, v := range hm {
		if v == 1 {
			res += k
		}
	}
	return res
}
