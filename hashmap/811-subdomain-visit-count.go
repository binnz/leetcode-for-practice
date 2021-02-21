package main

import "strings"
import "strconv"
import "fmt"

func subdomainVisits(cpdomains []string) []string {
	hm := make(map[string]int)
	for _, e := range cpdomains {
		s := strings.Split(e, " ")
		intVar, _ := strconv.Atoi(s[0])
		hm[s[1]] += intVar
		if idx := strings.Index(s[1], "."); idx != -1 {
			hm[s[1][idx+1:]] += intVar
			if i := strings.Index(s[1][idx+1:], "."); i != -1 {
				hm[s[1][idx+i+2:]] += intVar
			}
		}
	}
	res := []string{}
	for k, v := range hm {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
