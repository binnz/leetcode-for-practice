package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	hm := make(map[int][]int)
	row := len(mat)
	col := len(mat[0])
	for i := 0; i < row; i++ {
		l, r := 0, col
		for l < r {
			mid := l + (r-l)>>1
			if mat[i][mid] == 1 {
				l = mid + 1
			} else {
				r = mid
			}
		}

		hm[l] = append(hm[l], i)

	}
	vals := []int{}
	for k, _ := range hm {
		vals = append(vals, k)
	}
	sort.Ints(vals)
	res := []int{}
	for _, e := range vals {
		if k == 0 {
			break
		}
		if k < len(hm[e]) {
			res = append(res, hm[e][:k]...)
			k = 0
		} else {
			res = append(res, hm[e]...)
			k -= len(hm[e])
		}

	}
	return res
}
