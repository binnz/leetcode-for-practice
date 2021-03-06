package main

import "sort"

func findRightInterval(intervals [][]int) []int {
	hm := map[int]int{}
	for idx, e := range intervals {
		hm[e[0]] = idx
	}
	start := []int{}
	for k, _ := range hm {
		start = append(start, k)
	}
	sort.Ints(start)
	res := []int{}
	for _, e := range intervals {
		left, right := 0, len(intervals)
		for left < right {
			mid := left + (right-left)>>1
			if start[mid] == e[1] {
				left = mid
				break
			} else if start[mid] < e[1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == len(start) {
			res = append(res, -1)
		} else if start[left] == e[1] {
			res = append(res, hm[start[left]])
		} else if left == 0 {
			res = append(res, hm[start[0]])
		} else {
			res = append(res, hm[start[left]])
		}
	}
	return res
}
