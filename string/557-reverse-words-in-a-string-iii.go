package main

func reverseWords(s string) string {
	b := []byte(s)
	start := 0
	for i, v := range b {
		var end int
		if string(v) == " " || i == len(b)-1 {
			if i == len(b)-1 {
				end = i
			} else {
				end = i - 1
			}
			for start < end {
				b[start], b[end] = b[end], b[start]
				start++
				end--
			}
			start = i + 1
		}
	}
	return string(b)
}
