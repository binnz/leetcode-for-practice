package main

func distributeCandies(candyType []int) int {
	hm := map[int]int{}
	for _, e := range candyType {
		hm[e]++
	}
	n := len(candyType) / 2
	if len(hm) < n {
		return len(hm)
	}
	return n
}
