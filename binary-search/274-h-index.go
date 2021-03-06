package main

func hIndex(citations []int) int {
	l, r := 0, len(citations)+1
	for l < r {
		m := l + (r-l)>>1
		c := 0
		for _, e := range citations {
			if e >= m {
				c++
			}
		}
		if c >= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
