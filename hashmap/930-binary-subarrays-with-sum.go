package main

func numSubarraysWithSum(A []int, S int) int {
	hm := make(map[int]int)
	sum := 0
	res := 0
	hm[0] = 1
	for _, e := range A {
		sum += e
		if i, ok := hm[e-S]; ok {
			res += i
		}
		hm[sum]++
	}
	return res
}
