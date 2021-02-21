package main

import "strings"

func countCharacters(words []string, chars string) int {
	hm := make(map[byte]int)
	res := 0
	for i := 0; i < len(chars); i++ {
		hm[chars[i]]++
	}
	for _, e := range words {
		do := true
		for i := 0; i < len(e); i++ {
			if strings.Count(e, string(e[i])) > hm[e[i]] {
				do = false
				break
			}
		}
		if do {
			res += len(e)
		}

	}
	return res
}
