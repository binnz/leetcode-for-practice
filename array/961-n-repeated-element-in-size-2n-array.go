package main

func repeatedNTimes(A []int) int {
	hm := make(map[int]int)
	for _, e := range A {
		hm[e]++
		if hm[e] > 1 {
			return e
		}
	}
	return 0
}
