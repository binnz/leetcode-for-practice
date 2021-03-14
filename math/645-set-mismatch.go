package main

func findErrorNums(nums []int) []int {
	hm := map[int]int{}
	sum := 0
	dum := 0
	c := 0
	for i, e := range nums {
		sum += (i + 1)
		if hm[e] == 0 {
			c += e
		}
		hm[e]++
		if hm[e] == 2 {
			dum = e
		}
	}
	return []int{dum, sum - c}
}
