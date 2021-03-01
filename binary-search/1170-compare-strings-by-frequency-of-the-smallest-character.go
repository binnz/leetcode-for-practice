package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	qNum := make([]int, len(queries))
	wNum := make([]int, len(words))
	for idx, e := range queries {
		var char byte = 255
		var c int
		for i := 0; i < len(e); i++ {
			if e[i] < char {
				char = e[i]
				c = 1
			} else if e[i] == char {
				c++
			}
		}
		qNum[idx] = c
	}
	for idx, e := range words {
		var char byte = 255
		var c int
		for i := 0; i < len(e); i++ {
			if e[i] < char {
				char = e[i]
				c = 1
			} else if e[i] == char {
				c++
			}
		}
		wNum[idx] = c
	}
	res := make([]int, len(queries))
	sort.Ints(wNum)
	for idx, e := range qNum {
		left, right := 0, len(wNum)
		for left < right {
			mid := left + (right-left)>>1
			if wNum[mid] <= e {
				left = mid + 1
			} else {
				right = mid
			}
		}
		res[idx] = len(words) - left
	}
	return res
}
