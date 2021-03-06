package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	hm := map[int]int{}
	res := 0
	for _, i := range A {
		for _, j := range B {
			hm[i+j]++
		}
	}
	for _, x := range C {
		for _, y := range D {
			if c, ok := hm[-(x + y)]; ok {
				res += c
			}
		}
	}
	return res
}
