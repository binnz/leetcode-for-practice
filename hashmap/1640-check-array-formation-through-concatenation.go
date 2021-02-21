package main

func canFormArray(arr []int, pieces [][]int) bool {
	hm := map[int]int{}
	for idx, e := range pieces {
		hm[e[0]] = idx
	}
	i := 0
	for i < len(arr) {
		if idx, ok := hm[arr[i]]; ok {
			for _, v := range pieces[idx] {
				if arr[i] == v {
					i++
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}
