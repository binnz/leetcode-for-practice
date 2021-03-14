package main

func isPossible(nums []int) bool {
	hm := map[int]int{}
	for _, e := range nums {
		hm[e]++
	}
	tail := map[int]int{}
	for _, e := range nums {
		if hm[e] == 0 {
			continue
		}
		if tail[e-1] > 0 {
			tail[e-1]--
			tail[e]++
			hm[e]--
		} else if hm[e+1] > 0 && hm[e+2] > 0 {
			tail[e+2]++
			hm[e]--
			hm[e+1]--
			hm[e+2]--
		} else {
			return false
		}
	}
	return true
}
