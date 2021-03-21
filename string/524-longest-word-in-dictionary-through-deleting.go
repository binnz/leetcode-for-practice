package main

import "strings"

func findLongestWord(s string, dictionary []string) string {
	res := ""
	c := 0
	for _, str := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(str) {
			if s[i] == str[j] {
				j++
				i++
			} else {
				i++
			}
		}
		if j == len(str) {
			if j > c {
				res = str
				c = j
			} else if j == c {
				if strings.Compare(str, res) == -1 {
					res = str
				}
			}
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
