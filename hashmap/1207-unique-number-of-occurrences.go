package main

func uniqueOccurrences(arr []int) bool {
	var (
		hm        map[int]int      = map[int]int{}
		timeCount map[int]struct{} = map[int]struct{}{}
	)
	for _, e := range arr {
		hm[e]++
	}
	for _, v := range hm {
		if _, ok := timeCount[v]; ok {
			return false
		} else {
			timeCount[v] = struct{}{}
		}
	}
	return true
}
