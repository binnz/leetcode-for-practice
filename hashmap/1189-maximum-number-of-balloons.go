package main

func maxNumberOfBalloons(text string) int {
	hm := make(map[byte]int)
	for _, e := range text {
		hm[byte(e)]++
	}
	res := len(text)
	target := "balloon"
	for idx, v := range target {
		n, _ := hm[byte(v)]
		if n == 0 {
			return 0
		}
		if idx == 0 || idx == 1 || idx == 6 {
			if n < res {
				res = n
			}
		} else if idx == 2 || idx == 4 {
			if int(n/2) < res {
				res = int(n / 2)
			}
		} else {
			continue
		}
	}
	return res
}
