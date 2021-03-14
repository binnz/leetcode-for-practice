package main

import "sort"

func findLongestChain(pairs [][]int) int {
	res := 0
	arr := Pairs{}
	arr = pairs
	sort.Sort(arr)
	cur := -2 << 31
	for _, e := range arr {
		if cur < e[0] {
			res++
			cur = e[1]
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pairs [][]int

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	if p[i][1] == p[j][1] {
		return p[i][0] < p[j][0]
	}
	return p[i][1] < p[j][1]
}
func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
