package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	var res []string
	t := strings.Split(text, " ")
	for i := 0; i < len(t)-2; i++ {
		if t[i] == first && t[i+1] == second {
			res = append(res, t[i+2])
		}
	}
	return res
}
