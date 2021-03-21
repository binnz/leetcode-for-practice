package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	t := Time(timePoints)
	sort.Sort(t)
	res := 2 << 12
	for i := 0; i < len(t)-1; i++ {
		a := count(t[i], t[i+1])
		res = min(res, a)
	}
	a := count(t[len(t)-1], t[0])
	res = min(res, a)
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func atoi(s string) int {
	h, _ := strconv.Atoi(s[:2])
	m, _ := strconv.Atoi(s[3:])
	return h*60 + m
}

func count(a, b string) int {
	fmt.Println(b[:2])
	bc := atoi(b)
	ac := atoi(a)
	if bc >= ac {
		return bc - ac
	}
	return bc + 24*60 - ac
}

type Time []string

func (t Time) Len() int {
	return len(t)
}
func (t Time) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t Time) Less(i, j int) bool {
	s1 := t[i][:3]
	s2 := t[j][:3]
	if s1 < s2 {
		return true
	}
	if s1 == s2 {
		return t[i][3:] < t[j][3:]
	}
	return false
}
