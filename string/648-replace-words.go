package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	hm := map[string]struct{}{}
	for _, e := range dictionary {
		hm[e] = struct{}{}
	}
	res := []string{}
	words := strings.Split(sentence, " ")
	for _, e := range words {
		t := true
		for i := 1; i <= len(e); i++ {
			if _, ok := hm[string(e[:i])]; ok {
				res = append(res, e[:i])
				t = false
				break
			}
		}
		if t {
			res = append(res, e)
		}

	}
	return strings.Join(res, " ")
}
