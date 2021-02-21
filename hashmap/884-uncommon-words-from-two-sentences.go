package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	ha := make(map[string]int)
	hb := make(map[string]int)
	for _, e := range strings.Split(A, " ") {
		ha[e]++
	}
	for _, e := range strings.Split(B, " ") {
		hb[e]++
	}
	res := []string{}
	for k, v := range ha {
		_, ok := hb[k]
		if v == 1 && !ok {
			res = append(res, k)
		}
	}
	for k, v := range hb {
		_, ok := ha[k]
		if v == 1 && !ok {
			res = append(res, k)
		}
	}
	return res
}
