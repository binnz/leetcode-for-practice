package main

import "strings"
import "fmt"

func compressString(S string) string {
	var b strings.Builder
	if len(S) < 2 {
		return S
	}
	i, j := 0, 1
	c := 1
	for j < len(S) {
		if S[i] != S[j] {
			b.WriteByte(S[i])
			b.WriteString(fmt.Sprintf("%d", c))
			i = j
			j++
			c = 1
		} else {
			j++
			c++
		}
	}
	b.WriteByte(S[i])
	b.WriteString(fmt.Sprintf("%d", c))
	res := b.String()
	if len(res) >= len(S) {
		return S
	} else {
		return res
	}
}
