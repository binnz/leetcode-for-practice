package main

import (
	"strings"
)

func replaceSpaces(S string, length int) string {
	res := make([]string, length)
	for i := 0; i < length; i++ {
		if S[i] != 32 {
			res[i] = string(S[i])
		} else {
			res[i] = "%20"
		}
	}
	return strings.Join(res, "")
}
