package main

import "sort"

func findRadius(houses []int, heaters []int) int {
	res := []int{}
	sort.Ints(heaters)
	for _, e := range houses {
		left, right := 0, len(heaters)
		for left < right {
			mid := left + (right-left)>>1
			if heaters[mid] < e {
				left = mid + 1
			} else {
				right = mid
			}
		}
		t := 0
		if left == e {
			t = 0
		} else if left == 0 {
			t = heaters[0] - e
		} else if left == len(heaters) {
			t = e - heaters[len(heaters)-1]
		} else {
			t = min2(e-houses[left-1], houses[left]-e)
		}
		res = append(res, t)
	}
	var o int
	for _, e := range res {
		o = max2(o, e)
	}
	return o
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
