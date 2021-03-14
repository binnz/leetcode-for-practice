package main

import "math"

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		t := l*l + r*r
		if t == c {
			return true
		}
		if t < c {
			l++
		} else {
			r--
		}
	}
	return false
}
